package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Body   interface{} `json:"body"`
	Params string      `json:"params"`
}

func handleGetApi(w http.ResponseWriter, r *http.Request) {
	params := r.URL.RawQuery

	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{
		Body:   body,
		Params: params,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Server Accepting Body in GET request")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("GET /api", handleGetApi)

	fmt.Println("Listening on Port 3000")
	http.ListenAndServe(":3000", mux)
}
