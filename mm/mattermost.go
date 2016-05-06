package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	IconURL  string `json:"icon_url"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func Send(u string, p Payload) error {
	d, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	r := bytes.NewReader(d)
	_, err = http.Post(u, "application/json", r)
	if err != nil {
		return err
	}
	return nil
}
