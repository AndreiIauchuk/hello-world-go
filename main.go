package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App started!")
	http.HandleFunc("/", handleFunc())
	fmt.Println("App is listening 8080 port")
	http.ListenAndServe(":8080", nil)
}

func handleFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello world!"
		w.Write([]byte(msg))
		fmt.Println(msg)
	}
}
