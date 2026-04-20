package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/a-h/templ"
	"golang.org/x/time/rate"
)

// --- MOCK DATA STRUCTURES ---
type Topic struct {
	ID    string
	Title string
}

type Post struct {
	ID    string
	Title string
	Body  string
}

type Comment struct {
	ID     string
	Author string
	Text   string
	Points int
}

// Added CloudWord to track Live status and size
type CloudWord struct {
	Text   string
	IsLive bool
	Size   int 
}

// --- MOCK DATABASE ---
var topicsDB = make(map[string][]Topic)
var postsDB = make(map[string][]Post)
var commentsDB = []Comment{}





func init() {
	topicsDB["الوعي"] = []Topic{{"3", "The Edge of the Universe"}}
	topicsDB["الواقع"] = []Topic{{"1", "نقاشات عامه"}, {"2", "Ethics in the Modern Era"}}

	postsDB["1"] = []Post{
	
		{"p1", "هل فعلاً نعيش داخل محاكاة رقمية", "أفكر أحياناً أن كل شيء حولنا مجرد نظام مبرمج"},
		{"p2", "لماذا الوقت يمر بسرعة كلما كبرنا", "أشعر أن الأيام أصبحت أقصر مقارنة بطفولتي القديمة"},
		{"p3", "هل الصمت أبلغ من الكلام أحياناً", "في مواقف كثيرة الصمت يعبر أكثر من أي كلمات"},
		{"p4", "تجربة غريبة جعلتني أعيد التفكير بالحياة", "موقف بسيط غير نظرتي لكثير من الأمور اليومية"},
		{"p5", "هل نحن أحرار أم مجرد نتائج للظروف", "قراراتنا قد تكون نتيجة تراكمات لا نلاحظها أبداً"},
		{"p6", "لماذا نشعر بالوحدة رغم وجود الناس", "العزلة شعور داخلي لا علاقة له بالمحيط الخارجي"},
		{"p7", "هل الذكاء الاصطناعي سيغير مستقبل البشر", "التطور الحالي يوحي بقفزة كبيرة في طريقة عيشنا"},
		{"p8", "ما معنى أن تكون إنساناً في هذا العصر", "الهوية أصبحت أكثر تعقيداً مع تغير القيم المستمر"},
		{"p9", "هل الأحلام رسائل أم مجرد نشاط عشوائي", "بعض الأحلام تبدو واقعية بشكل يصعب تجاهله"},
		{"p10", "لماذا نخاف من المجهول رغم فضولنا", "الخوف والفضول يتعايشان داخلنا بشكل غريب دائماً"},

	}


	postsDB["2"] = []Post{
	
	
		{"p21", "هل العقل أم القلب يقود قراراتنا فعلاً", "في الواقع كلاهما يتداخلان بطرق معقدة للغاية"},
		{"p22", "لماذا نشعر بالندم بعد اتخاذ بعض القرارات", "التفكير الزائد بعد الفعل يخلق هذا الشعور دائماً"},
		{"p23", "هل يمكن للإنسان أن يتغير جذرياً بمرور الوقت", "التجارب القوية قادرة على إعادة تشكيل الشخصية"},
		{"p24", "لماذا ننجذب للأشياء الغامضة بطبيعتنا البشرية", "الغموض يحفز الفضول ويجعلنا نبحث عن الإجابات"},
		{"p25", "هل هناك معنى ثابت للحياة أم يتغير دائماً", "المعنى يختلف حسب التجارب والظروف الشخصية لكل فرد"},
		{"p26", "كيف تؤثر البيئة على طريقة تفكير الإنسان", "المحيط يلعب دوراً كبيراً في تشكيل القناعات"},
		{"p27", "هل يمكن أن يكون الفشل خطوة ضرورية للنجاح", "الفشل غالباً يحمل دروساً لا يوفرها النجاح بسهولة"},
		{"p28", "لماذا نميل لتأجيل الأمور المهمة في حياتنا", "التسويف مرتبط بالخوف أو بعدم وضوح الهدف الحقيقي"},
		{"p29", "هل يمكن تحقيق التوازن بين العمل والحياة الشخصية", "الأمر يتطلب وعي مستمر وتحديد أولويات واضحة دائماً"},
		{"p30", "ما الذي يجعل بعض الأشخاص أكثر إبداعاً من غيرهم", "الاختلاف في التفكير والتجربة يصنع هذا التميز"},

		
	}

	postsDB["3"] = []Post{
	
		{"p11", "هل النجاح يعتمد على الحظ أم الجهد", "أرى أن التوازن بينهما هو العامل الحقيقي للنتائج"},
		{"p12", "كيف تؤثر التكنولوجيا على علاقاتنا اليومية", "التواصل أصبح أسهل لكن العمق أصبح أقل بكثير"},
		{"p13", "هل يمكن للإنسان أن يفهم نفسه بالكامل", "كلما تعلمت أكثر أكتشف أنني لا أعرف نفسي"},
		{"p14", "لماذا نكرر نفس الأخطاء رغم إدراكنا لها", "العادات أقوى من الوعي في كثير من الأحيان"},
		{"p15", "هل السعادة هدف أم نتيجة لأسلوب الحياة", "أعتقد أنها نتيجة تراكم قرارات صغيرة يومية"},
		{"p16", "ما الفرق بين المعرفة والحكمة في حياتنا", "المعرفة معلومات بينما الحكمة استخدام صحيح لها"},
		{"p17", "هل الماضي يؤثر علينا أكثر مما نعتقد", "ذكريات قديمة تتحكم بسلوكنا دون أن نلاحظها"},
		{"p18", "لماذا يصعب اتخاذ القرارات المصيرية أحياناً", "الخوف من الخطأ يمنعنا من التقدم بثقة كاملة"},
		{"p19", "هل يمكن أن نعيش بدون مقارنة أنفسنا بالآخرين", "المقارنة عادة متجذرة يصعب التخلص منها بسهولة"},
		{"p20", "ما الذي يجعل فكرة ما تبدو مقنعة لنا", "العاطفة أحياناً تتغلب على المنطق في الحكم"},

			}

}

// Helper: Generate 100+ Words
func getWords() []CloudWord {
	words := []string{
		"الواقع",
		"الوعي",
		"الزمن",
		"الذكاء",
		"الكون",
		"الفكر",
		"الهوية",
	}

	rand.Seed(time.Now().UnixNano())

	var cloud []CloudWord
	for i, w := range words {
		cloud = append(cloud, CloudWord{
			Text:   w,
			IsLive: rand.Float32() < 0.3,
			Size:   (i % 3) + 1,
		})
	}
	return cloud
}


var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	limiter, exists := visitors[ip]
	if !exists {
		// Allow 5 requests per second, with a burst of 10
		limiter = rate.NewLimiter(5, 10)
		visitors[ip] = limiter
	}
	return limiter
}

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		limiter := getVisitor(ip)
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ==========================================
// 2. SECURITY HEADERS MIDDLEWARE
// ==========================================
func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent Clickjacking
		w.Header().Set("X-Frame-Options", "DENY")
		// Prevent MIME-sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")
		// XSS Protection
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		// Strict Transport Security (Force HTTPS)
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		
		next.ServeHTTP(w, r)
	})
}


func AnalyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler (process the request)
		next.ServeHTTP(w, r)

		duration := time.Since(startTime)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		
		// Here you calculate enters and behavior. 
		// In production, save this to a Database (Postgres) or Redis instead of printing.
		log.Printf("[STAT] IP: %s | Method: %s | Path: %s | Duration: %v | User-Agent: %s\n", 
			ip, r.Method, r.URL.Path, duration, r.UserAgent())
	})
}


func main() {

	port := os.Getenv("PORT")
if port == "" {
	port = "1888"
}



	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(HomePage(getWords())).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/topics", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(150 * time.Millisecond) 
		word := r.URL.Query().Get("word")
		topics := topicsDB[word]
		if len(topics) == 0 { topics = topicsDB["default"] }
		templ.Handler(TopicsList(word, topics)).ServeHTTP(w, r)
	})

	mux.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		topicID := r.URL.Query().Get("id")
		posts := postsDB[topicID]
		templ.Handler(PostsList(posts)).ServeHTTP(w, r)
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


	handler := AnalyticsMiddleware(SecurityMiddleware(RateLimitMiddleware(mux)))

	// Security: Use a custom server struct to enforce timeouts (Prevents Slowloris DDoS attacks)
	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,  // Max time to read request
		WriteTimeout: 10 * time.Second, // Max time to write response
		IdleTimeout:  15 * time.Second, // Max time to keep connection alive
	}

	fmt.Println("Smart Base Server running on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
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
