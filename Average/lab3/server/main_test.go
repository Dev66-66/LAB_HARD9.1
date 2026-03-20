package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello from Lab 3 Go Server!"
	body, _ := io.ReadAll(rr.Body)
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}
