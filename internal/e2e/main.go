package main

import "github.com/varkey98/go-re2"

func main() {
	re := re2.MustCompile("hello")

	if !re.MatchString("hello world") {
		panic("failed to match")
	}
}
