package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", MessengerVerify)

	fmt.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MessengerVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		challenge := r.URL.Query().Get("hub.challenge")
		verify_token := r.URL.Query().Get("hub.verify_token")

		if len(verify_token) > 0 && len(challenge) > 0 && verify_token == "developers-are-gods" {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, challenge)
			return
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)
	fmt.Fprintf(w, "Bad Request")
}
