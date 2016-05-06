package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Print("ok")
	}))
	defer server.Close()

	p := Payload{
		IconURL:  "http://example.com/image/icon",
		Username: "user",
		Text:     "test",
	}
	err := Send(server.URL, p)
	if err != nil {
		t.Error(err)
	}
}
