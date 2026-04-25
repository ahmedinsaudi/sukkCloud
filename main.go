package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

type Topic struct {
	ID    string
	Title string
}

type Source struct {
	Word  string `json:"word"`
	Title string `json:"sourceTitle"`
}

type AudienceData struct {
	Origins   string
	Interests string
}

type Post struct {
	ID            string
	Title         string
	Body          string
	Sources       []Source
	Views         int
	Resonances    int
	CommentsCount int
	Audience      AudienceData
}

type Comment struct {
	ID     string
	Author string
	Text   string
	Points int
}

type CloudWord struct {
	Text   string
	IsLive bool
	Size   int
}

type Notification struct {
	ID      string
	Title   string
	Message string
	PostID  string
	Time    string
}

// --- MOCK DATABASES ---
var topicsDB = make(map[string][]Topic)
var postsDB = make(map[string][]Post)
var commentsDB = []Comment{}
var savedPostsDB = []Post{}
var notificationsDB = []Notification{}

func init() {
	// INTENTIONALLY EMPTY FOR TESTING
	
	initPost()

	savedPostsDB = []Post{
		{ID: "sp1", Title: "The Fundamentals of Stoicism", Body: "A deep dive into emotional resilience... "},
		{ID: "sp2", Title: "Quantum Entanglement Explained", Body: "How particles communicate across the void..."},
	}


	notificationsDB = append([]Notification{
		{ID: "n_555" , Title: "New Resonance Event", Message: "A mind published a thought connected to your path: " },
	}, notificationsDB...)

	
}

func getWords() []CloudWord {
	baseWords := []string{
		"الواقع",
		"الوعي",
		"الزمن",
		"الذكاء",
		"الكون",
		"الفكر",
		"الهوية",
		"الكون",
		"الفكر",
		"الهوية",
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(baseWords), func(i, j int) { baseWords[i], baseWords[j] = baseWords[j], baseWords[i] })

	var cloud []CloudWord
	for i, w := range baseWords {
		size := 3
		if i%4 == 0 { size = 1 } else if i%2 == 0 { size = 2 }
		cloud = append(cloud, CloudWord{Text: w, IsLive: rand.Float32() < 0.15, Size: size})
	}
	return cloud
}



func main() {
	mux := http.ServeMux{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(HomePage(getWords())).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/topics", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(150 * time.Millisecond) // Simulate DB delay
		word := r.URL.Query().Get("word")
		topics, exists := topicsDB[word]
		if !exists { topics = topicsDB["default"] }
		templ.Handler(TopicsList(word, topics)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(150 * time.Millisecond)
		topicID := r.URL.Query().Get("id")
		posts := postsDB[topicID]
		templ.Handler(PostsList(posts)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/post_single", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var foundPost Post
		for _, p := range postsDB["1"] {
			if p.ID == id { foundPost = p; break }
		}
		templ.Handler(PostsList([]Post{foundPost})).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/comments", func(w http.ResponseWriter, r *http.Request) {

		var commentsDB = []Comment{}
		rand.Seed(time.Now().UnixNano())
		for i := 1; i <= 2; i++ {
			commentsDB = append(commentsDB, Comment{
				ID:     fmt.Sprintf("c%d", i),
				Author: authors[rand.Intn(len(authors))],
			Text:   arabicComments[rand.Intn(len(arabicComments))],
				Points: 50 + (i * 25),
			})
		}

		time.Sleep(150 * time.Millisecond)
		pageStr := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(pageStr)
		if page == 0 { page = 1 }

		start := (page - 1) * 5
		end := start + 5
		if end > len(commentsDB) { end = len(commentsDB) }

		displayed := []Comment{}
		if start < len(commentsDB) { displayed = commentsDB[start:end] }
		hasMore := end < len(commentsDB)

		templ.Handler(CommentsSection(displayed, page, hasMore, 450)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/saved", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(SavedPostsList(savedPostsDB)).ServeHTTP(w, r)
	})

	// FIXED: Cache Control to ensure new notifications are fetched
	mux.HandleFunc("/api/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		templ.Handler(NotificationsList(notificationsDB)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/publish", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		text := r.FormValue("text")
		// notify := r.FormValue("notify") == "on"
		branchFrom := r.FormValue("branchFrom") // New: Branch context
		sourcesJSON := r.FormValue("sources")

		var sources []Source
		if sourcesJSON != "" { json.Unmarshal([]byte(sourcesJSON), &sources) }

		title := "A New Origin Thought"
		if branchFrom != "" {
			title = "Branched from: " + branchFrom
		}

		newID := "p_new_" + strconv.Itoa(rand.Intn(10000))
		newPost := Post{
			ID: newID, Title: title, Body: text, Sources: sources,
			Views: 1, Resonances: 0, CommentsCount: 0, Audience: AudienceData{Origins: "Just you", Interests: "The Unknown"},
		}

		postsDB["1"] = append([]Post{newPost}, postsDB["1"]...)

		if true {
			notificationsDB = append([]Notification{
				{ID: "n_" + newID, Title: "New Resonance Event", Message: "A mind published a thought connected to your path: " + title, PostID: newID, Time: "Just now"},
			}, notificationsDB...)
			w.Header().Set("HX-Trigger", "newNotification")
		}

		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Mobile-First Advanced Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", &mux))
}



var authors = []string{
	"Ahmed Alqahtani",
	"Mohammed Ali",
	"Fahad Alshammari",
	"Abdullah Khan",
	"Yousef Hassan",
	"Khalid Mansour",
	"Nora Alotaibi",
	"Lina Ahmed",
}

var arabicComments = []string{
	"وجهة نظر جميلة فعلاً، أوافقك جزئياً.",
	"أعتقد أن الموضوع أعمق مما يبدو.",
	"تحليل منطقي لكن يحتاج أمثلة أكثر.",
	"هذا نفس اللي لاحظته في حياتي اليومية.",
	"صراحة الفكرة مثيرة للاهتمام.",
	"لا أتفق بالكامل، لكن طرحك جيد.",
	"أشوف أن الواقع مختلف حسب التجربة.",
	"كلامك فتح لي زاوية تفكير جديدة.",
}
