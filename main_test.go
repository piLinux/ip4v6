package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.RemoteAddr = "127.0.0.1"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HTTPHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code got: %v want %v",
			status, http.StatusOK)
	}

	expected := "127.0.0.1\n"
	if rr.Body.String() != expected {
		t.Errorf("Unexpected body got: %v want %v",
			rr.Body.String(), expected)
	}
}
