package main

import (
	//"fmt"
	"log"
	"net/http"
	"portfolio/handlers"
)

func main() {

	// Define HTTP handlers
	http.HandleFunc("/", handlers.Index)
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Defining server protocol and http port
	url := "http://localhost:8080"
	log.Println("Server is running on", url)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}