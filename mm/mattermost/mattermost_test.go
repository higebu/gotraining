package mattermost

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

	m, err := New(server.URL)

	p := Payload{
		IconURL:  "http://example.com/image/icon",
		Username: "user",
		Text:     "test",
	}
	err = m.Send(p)
	if err != nil {
		t.Error(err)
	}
}
