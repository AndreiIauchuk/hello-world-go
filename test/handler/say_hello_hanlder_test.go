package test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"hello-world-go/internal/handler"
)

func TestSayHelloRequestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/say-hello", nil)
    w := httptest.NewRecorder()
    handler.SayHelloRequestHandler(w, req)
    
	if status := w.Code; status != http.StatusOK {
        t.Errorf(
			"handler returned wrong status code: got %v want %v",
            status,
			http.StatusOK,
		)
    }

    if contentType := w.Header().Get("Content-Type"); contentType != "application/json" {
        t.Errorf(
			"handler returned wrong content type: got %v want %v",
		 	contentType,
			"application/json",
		)
    }

    var res handler.ApiResponse
    if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
        t.Errorf("error decoding response body: %v", err)
    }

    expectedMessage := "Hello from Golang!"
    if res.Text != expectedMessage {
        t.Errorf(
			"handler returned unexpected body: got %v want %v",
            res.Text,
			expectedMessage,
		)
    }
}