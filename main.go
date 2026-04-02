package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

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

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := SuccessResponse{
		Message: "Got signup endpoint",
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/", notFoundHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
