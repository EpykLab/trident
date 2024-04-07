package webserver

import (
	"html/template"
	"log"
	"net/http"

	logger "github.com/Epyklab/trident/utils/logging"
)

func Router() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/search", searchHandler)

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Log the login attempt
		logger.Log.Printf("Login attempt. Username: %s. Password: %s", r.FormValue("username"),
			r.FormValue("password"))
	}
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Log the search query
		logger.Log.Printf("Search query: %s\n", r.URL.Query().Get("q"))
	}
	tmpl := template.Must(template.ParseFiles("templates/search.html"))
	tmpl.Execute(w, nil)
}
