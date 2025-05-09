package main

import (
	"net/http"
	"time"
)

type PingResult struct {
	URL     string  `json:"url"`
	Up      bool    `json:"up"`
	Latency float64 `json:"latency"` // in ms
}

func pingURLs(urls []string) []PingResult {
	results := make([]PingResult, len(urls))
	ch := make(chan PingResult)

	for _, url := range urls {
		go func(u string) {
			start := time.Now()
			resp, err := http.Get(u)
			latency := time.Since(start).Seconds() * 1000

			if err != nil || resp.StatusCode >= 400 {
				ch <- PingResult{URL: u, Up: false, Latency: latency}
			} else {
				ch <- PingResult{URL: u, Up: true, Latency: latency}
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		results[i] = <-ch
	}

	return results
}
