package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	// START OMIT
	type Tag struct {
		Tag    string  `json:"tag"`
		Weight float32 `json:"weight"`
	}
	type InfoResponse struct {
		Name string `json:"name"`
		Tags []*Tag `json:"tags"`
	}
	http.HandleFunc("/api/v1/info", func(w http.ResponseWriter, r *http.Request) {
		info := &InfoResponse{
			Name: "CEWIT",
			Tags: []*Tag{
				{"Go", 1.0}, {"HTTP", 0.64}, {"JSON", 0.4},
			},
		}
		b, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.Write(b)
	})
	// END OMIT
	log.Println("Starting server on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
