package main

import (
	"log"
	"net/http"
	"url_shortener/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/redirect/", handlers.RedirectURL)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
