package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Index handles requests to "/" and "/Home"
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/home" {
		if r.Method == http.MethodGet {
			serveTemplate(w, "index.html")
		}
	}else{
		if r.Method == http.MethodGet {
			http.NotFound(w, r)
		}
	}
}


// serveTemplate loads and executes a template file
func serveTemplate(w http.ResponseWriter, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		http.Error(w, "408 Page not Found", http.StatusNotFound)
		return
	}
	tmpl := template.Must(template.ParseFiles(filename))
	errr := tmpl.Execute(w, nil)
	if errr != nil {
		log.Println("500 Internal Server Error")
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
