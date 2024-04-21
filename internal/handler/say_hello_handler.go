package handler

import (
	"net/http"
	"encoding/json"
)

type ApiResponse struct {
	Text string `json:"text"`
}

func HandleSayHello() {
    http.HandleFunc("/say-hello", SayHelloRequestHandler)
}

func SayHelloRequestHandler(w http.ResponseWriter, r *http.Request) {
    res := ApiResponse{Text: "Hello from Golang!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}