package main

import (
	"flag"
	"fmt"
	"os"
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
	p := Payload{
		IconURL:  *iconURL,
		Username: *username,
		Text:     *text,
	}
	err := Send(*url, p)
	if err != nil {
		fmt.Println(err)
	}
}
