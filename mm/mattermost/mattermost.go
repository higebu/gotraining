package mattermost

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Payload struct {
	IconURL  string `json:"icon_url"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

type Mattermost struct {
	URL string
}

func New(u string) (Mattermost, error) {
	pu, _ := url.Parse(u)
	if pu.Scheme != "http" && pu.Scheme != "https" {
		return Mattermost{}, errors.New(fmt.Sprintf("scheme (%s) must be http or https", pu.Scheme))
	}
	return Mattermost{URL: u}, nil
}

func (m *Mattermost) Send(p Payload) error {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	r := bytes.NewReader(d)
	_, err = http.Post(m.URL, "application/json", r)
	if err != nil {
		return err
	}
	return nil
}
