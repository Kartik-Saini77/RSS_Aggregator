package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, _ *http.Request) {
	respondWithJson(w, 400, struct{}{})
}
