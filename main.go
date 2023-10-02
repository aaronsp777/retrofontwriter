package main

import (
	"fmt"
	"unicode/utf8"
)

var message = []string{
	"Foo Basic",
	"£π BYTES FREE",
	"",
	"READY",
}

func SizeMessage(m []string) (rows, cols int) {
	rows = len(message)
	for _, r := range m {
		c := utf8.RuneCountInString(r)
		if cols < c {
			cols = c
		}
	}
	return
}

func main() {
	rows, cols := SizeMessage(message)
	fmt.Println(rows, cols)
}
