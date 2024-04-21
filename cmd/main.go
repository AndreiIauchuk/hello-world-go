package main

import (
	"fmt"
	"net/http"
	handler "hello-world-go/internal/handler"
)

func main() {
	fmt.Println("App started!")
	handler.HandleSayHello()
	fmt.Println("App is listening 8080 port")
	http.ListenAndServe(":8080", nil)
}
