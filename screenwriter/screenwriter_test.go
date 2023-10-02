package screenwriter

import (
	"fmt"
	"testing"
)

func TestAscii(t *testing.T) {
	tests := map[rune]int{
		'A': 1,
		'Z': 26,
		' ': 32,
		'a': 1 + 256,
		'z': 26 + 256,
	}
	for in, want := range tests {
		got := Ascii(in).offset / 8
		if got != want {
			t.Errorf("Ascii(%d) = %d, want %d", in, got, want)
		}
	}
}

func TestPixelAt(t *testing.T) {
	tests := []rune{'A', 'a', 'Z', 'z', 'π', '£', '='}
	for _, in := range tests {
		fmt.Printf("\nchar: %q\n", in)
		c := Ascii(in)
		for y := 0; y < 8; y++ {
			fmt.Print(":")
			for x := 0; x < 8; x++ {
				if c.pixelAt(x, y) {
					fmt.Print("X")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
