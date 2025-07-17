package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Error: %v\t%v", code, msg)
	}

	type errResponse struct {
		Error string `json:"error"`
		Time time.Time `json:"time"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
		Time: time.Now(),
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
