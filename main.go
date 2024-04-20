package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App started!")
    http.HandleFunc("/say-hello", handleRequestFunc())
	fmt.Println("App is listening 8080 port")
	http.ListenAndServe(":8080", nil)
}

type ApiResponse struct {
	Text string `json:"text"`
}

func handleRequestFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
        response := ApiResponse{Text: "Hello from Golang!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
