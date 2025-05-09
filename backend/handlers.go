package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var latestResults []PingResult

func handlePing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var urls []string
	err = json.Unmarshal(body, &urls)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	results := pingURLs(urls)
	latestResults = results

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func handleResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if latestResults == nil {
		json.NewEncoder(w).Encode([]PingResult{}) // send empty array
		return
	}
	json.NewEncoder(w).Encode(latestResults)
}
