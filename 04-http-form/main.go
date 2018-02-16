package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// START OMIT
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		count := r.FormValue("count")
		fmt.Fprintf(w, "%s widgets have been ordered!\n", count)
	})
	// Assumes we are running from the repo root
	fs := http.Dir("04-http-form/static")
	http.Handle("/", http.FileServer(fs))
	// END OMIT
	log.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
