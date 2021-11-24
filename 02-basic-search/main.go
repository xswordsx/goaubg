package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	log.Println("Listening on http://localhost:8888/search")

	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}

type result struct {
	Title   string `json:"title,omitempty"`
	Address string `json:"address,omitempty"`
}

// handleSearch searches Google for the "Golang" keyword and
// returns the top results for Text, Image and Video.
func handleSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling route", r.URL)

	results := []result{textSearch(), imgSearch(), videoSearch()}
	enc := json.NewEncoder(w)
	if r.URL.Query().Get("pretty") == "true" {
		enc.SetIndent("", "\t")
	}
	enc.Encode(results)
}

func textSearch() result {
	return result{"The Go Programming Language", "https://golang.org"}
}

func imgSearch() result {
	return result{"The Go Gopher", "https://miro.medium.com/max/1400/0*7vQ8eRc28yz9k__r.png"}
}

func videoSearch() result {
	return result{"Golang Tutorial #1 - An Introduction to Go Programming", "https://www.youtube.com/watch?v=75lJDVT1h0s"}
}
