package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/higebu/gotraining/mm/mattermost"
)

func main() {
	url := flag.String("url", "", "Incoming webhook url")
	iconURL := flag.String("icon-url", "", "")
	username := flag.String("username", "bot", "")
	text := flag.String("text", "Hello, this is some text.\nThis is more text.", "")
	flag.Parse()
	if *url == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	m, err := mattermost.New(*url)
	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
		os.Exit(0)
	}

	p := mattermost.Payload{
		IconURL:  *iconURL,
		Username: *username,
		Text:     *text,
	}
	err = m.Send(p)
	if err != nil {
		fmt.Println(err)
	}
}
