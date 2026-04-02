package main

import (
	"encoding/json"
	"go-auth/api/handler/signIn"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := ErrorResponse{Error: "Page was not found at all"}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(data)
}

func main() {
	http.HandleFunc("/signup", signIn.Post)
	http.HandleFunc("/", notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
