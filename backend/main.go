package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/results", handleResults)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
