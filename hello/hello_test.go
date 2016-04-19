package hello

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input    string
	expected string
}{
	{"higebu", "Hello, higebu"},
	{"hitsumabushi", "Hello, hitsumabushi"},
}

func TestHello(t *testing.T) {
	for _, v := range tests {
		if Hello(v.input) != v.expected {
			t.Error(v)
		}
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("higebu")
	}
}

func ExampleHello() {
	fmt.Println(Hello("higebu"))
	// Output: Hello, higebu
}
