package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.Int("int", 1, "intですよ")
	s := flag.String("string", "str", "stringですよ")
	flag.Parse()
	fmt.Printf("%d, %s\n", *i, *s)
}
