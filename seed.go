
package main


func initPost() {

	topicsDB["الوعي"] = []Topic{{"3", "العلم بالاشياء المهمه ... طريق عظيم"}}
	topicsDB["الواقع"] = []Topic{{"1", "نقاشات عامه"}, {"2", "Ethics in the Modern Era"}}


	// ------------------ 1 ------------------
	postsDB["1"] = append(postsDB["1"], Post{
		ID:    "p1",
		Title: "هل نعيش داخل محاكاة رقمية",
		Body:  "أفكر أحياناً أن العالم كله مجرد نظام مبرمج",
		Views: 1450, Resonances: 320, CommentsCount: 41,
		Audience: AudienceData{
			Origins:   "مجتمعات التقنية والفلسفة",
			Interests: "الذكاء الاصطناعي، الوعي",
		},
	})

	postsDB["1"] = append(postsDB["1"], Post{
		ID:    "p2",
		Title: "لماذا يمر الوقت بسرعة",
		Body:  "كلما كبرنا نشعر أن الأيام تختصر نفسها",
		Views: 980, Resonances: 210, CommentsCount: 33,
		Audience: AudienceData{
			Origins:   "مهتمون بعلم النفس",
			Interests: "الإدراك، الزمن",
		},
	})

	postsDB["1"] = append(postsDB["1"], Post{
		ID:    "p3",
		Title: "هل الصمت أقوى من الكلام",
		Body:  "في بعض اللحظات الصمت يحمل معاني أعمق",
		Views: 760, Resonances: 150, CommentsCount: 20,
		Audience: AudienceData{
			Origins:   "مجتمعات أدبية",
			Interests: "التعبير، المشاعر",
		},
	})

	// ------------------ 2 ------------------
	postsDB["2"] = append(postsDB["2"], Post{
		ID:    "p21",
		Title: "العقل أم القلب",
		Body:  "قراراتنا نتيجة تفاعل معقد بين المنطق والعاطفة",
		Views: 1320, Resonances: 410, CommentsCount: 65,
		Audience: AudienceData{
			Origins:   "مفكرون وباحثون",
			Interests: "الوعي، اتخاذ القرار",
		},
	})

	postsDB["2"] = append(postsDB["2"], Post{
		ID:    "p22",
		Title: "لماذا نشعر بالندم",
		Body:  "التفكير بعد القرار يخلق سيناريوهات بديلة",
		Views: 870, Resonances: 190, CommentsCount: 29,
		Audience: AudienceData{
			Origins:   "علم النفس",
			Interests: "السلوك، الإدراك",
		},
	})

	postsDB["2"] = append(postsDB["2"], Post{
		ID:    "p23",
		Title: "هل يمكن أن نتغير",
		Body:  "التجارب القاسية تعيد تشكيل الإنسان بالكامل",
		Views: 1110, Resonances: 275, CommentsCount: 48,
		Audience: AudienceData{
			Origins:   "تطوير الذات",
			Interests: "النمو، التغيير",
		},
	})

	// ------------------ 3 ------------------
	postsDB["3"] = append(postsDB["3"], Post{
		ID:    "p11",
		Title: "هل النجاح حظ أم جهد",
		Body:  "النجاح غالباً نتيجة مزيج بين الاثنين",
		Views: 1560, Resonances: 500, CommentsCount: 72,
		Audience: AudienceData{
			Origins:   "رواد أعمال",
			Interests: "النجاح، الإنتاجية",
		},
	})

	postsDB["3"] = append(postsDB["3"], Post{
		ID:    "p12",
		Title: "تأثير التكنولوجيا على العلاقات",
		Body:  "قربتنا رقمياً لكنها قللت العمق الإنساني",
		Views: 940, Resonances: 230, CommentsCount: 37,
		Audience: AudienceData{
			Origins:   "مجتمع رقمي",
			Interests: "التواصل، التقنية",
		},
	})

	postsDB["3"] = append(postsDB["3"], Post{
		ID:    "p13",
		Title: "هل نفهم أنفسنا حقاً",
		Body:  "كلما تعلمنا أكثر ندرك جهلنا بأنفسنا",
		Views: 1020, Resonances: 260, CommentsCount: 44,
		Audience: AudienceData{
			Origins:   "فلسفة",
			Interests: "الذات، الوعي",
		},
	})

}