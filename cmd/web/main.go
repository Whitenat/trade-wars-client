package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	//Use the http.NewServeMux() function to initialize a new servemux
	//then register the home function as the handler for the "/" URL pattern
	//Finally, register the 2 new handler functions and URL patterns
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/navigation", showNavigation)
	mux.HandleFunc("/trade", showTrade)
	mux.HandleFunc("/chat", showChat)
	//mux.HandleFunc("/snippet/create", createSnippet)

	godotenv.Load()

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", serveTemplate)

	//Use the http.ListenAndServe() function to start a new web server
	//We pass in 2 params: the TCP network address to liston on, and the
	//servemux we just created. If http.ListenAndServe() returns an error
	//that any error return by http.ListenAndServe() is always non-nil
	PORT := os.Getenv("PORT")
	log.Println("Starting server on")
	err := http.ListenAndServe(":" + PORT, mux)
	log.Fatal(err)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
