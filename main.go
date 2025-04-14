package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"zenn-rss/models"
)

func main() {
	http.HandleFunc("/articles", articlesHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://zenn.dev/urakawa_jinsei/feed?all=1")
	if err != nil {
		http.Error(w, "Failed to fetch feed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var rss models.RSS
	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		http.Error(w, "Failed to parse feed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(rss.Channel.Items); err != nil {
		http.Error(w, "Failed to encode items to JSON", http.StatusInternalServerError)
	}
}
