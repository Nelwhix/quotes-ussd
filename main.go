package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

type AfricaTalkingCallback struct {
	SessionId   string `json:"sessionId"`
	ServiceCode string `json:"serviceCode"`
	Text        string `json:"text"`
}

func main() {
	http.HandleFunc("/quote", serveQuotes)
	log.Printf("Server starting on port %v", PORT)

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatalf("error occured starting server: %v", err.Error())
	}
}

func serveQuotes(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type. Only JSON data is allowed.", http.StatusUnsupportedMediaType)
		return
	}

	var request AfricaTalkingCallback
	var response string

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusInternalServerError)
		return
	}

	if request.Text == "" {
		response = "Welcome to Quotsy, I give you inspiring quotes depending on the mood you are select\n"
		response += "1. Enter 1 to continue"
	} else if request.Text == "1" {
		response = "You pressed 1!"
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = fmt.Fprint(w, response)

	if err != nil {
		http.Error(w, "Failed to format response", http.StatusInternalServerError)
		return
	}
}
