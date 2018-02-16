package main

import (
	"log"
	"net/http"
)

func main() {
	// Assumes we are running from the repo root
	fs := http.Dir("03-http-html/static")
	http.Handle("/", http.FileServer(fs))
	log.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
