// Package screenwriter writes chars from a font to a bitmap.
package screenwriter

import (
	_ "embed"
)

//go:embed chargen.rom
var chargen []byte

type Char struct {
	offset int
}

func screenCode(p int) Char {
	return Char{offset: p * 8}
}

// Behavior gleaned by experiments in Basic V2
// char       offset
//  0         128  // Control characters (inverse letters)
//  13        19   // CR ignores quote mode
//  14        128  // Control characters (inverse letters)
//  20        12   // DEL
//  21        128  // Control characters (inverse letters)
//  32        0   // typewriter symbols
//  64       -64  // letters
//  96       -32  // Drawing symbos
//  128       64
//  141      -109  // Shift CR ignores quote mode
//  142       64
//  160      -64
//  192      -128
//  255      -161  // pi symbol
// Generated by:
// 5 dimo(255)
// 10 fori=0to255
// 12 poke212,0
// 15 printchr$(147);: rem clear screen
// 20 poke212,1:rem quote mode
// 30 printchr$(i)
// 40 o(i)=peek(1024)-i
// 50 next
// 55 l=0
// 59 poke212,0
// 60 fori=0to255
// 65 ifo(i)<>l thenprinti,o(i)
// 70 l=o(i)
// 75 next

func Petscii(r rune, shifted bool) Char {
	p := int(r)
	if shifted {
		p |= 256
	}
	switch {
	case 0 <= r && r <= 31: // Control chars
		return screenCode(p + 128) // Inverse Letters
	case 32 <= r && r <= 63: // Punctuation & Numbers
		return screenCode(p)
	case 64 <= r && r <= 95: // Uppercase / LowerCase
		return screenCode(p - 64)
	case 96 <= r && r <= 127: // Drawing Symbols / Uppercase
		return screenCode(p - 32)
	case 128 <= r && r <= 159: // Shifted Control chars
		return screenCode(p + 64)
	case 160 <= r && r <= 191: // Block characters
		return screenCode(p - 64)
	case 192 <= r && r <= 254: // Repeate of Drawing Symbols
		return screenCode(p - 128)
	case r == 255: // Pi Symbol
		return screenCode(94)
	default: // Unknown rune
		return screenCode(416) // Shifted space
	}
}

func Ascii(r rune) Char {
	switch {
	case r == '_':
		return Petscii(164, false) // Line draw: underscore
	case r == '\\':
		return Petscii(109, false) // Line draw: backslash
	case r == '`':
		return Petscii(39, false) // Closest thing is forward tick
	case 32 <= r && r <= 95:
		// Note: ^ maps to Up Arrow
		return Petscii(r, false)
	case 96 <= r && r <= 127:
		// Known bugs:
		// { maps to [
		// | maps to British Pound £
		// } maps to ]
		// ~ maps to Up Arrow
		// DEL maps to Left Arrow
		return Petscii(r-32, true)
	case r == '£':
		return Petscii(92, false)
	case r == 'π':
		return Petscii(126, false)
	case r == '~':
		return Petscii(168, false)
	default:
		return Petscii(126, true) // Checkerboard
	}
}

func (c *Char) pixelAt(x, y int) bool {
	v := chargen[c.offset+y]
	v = v << x
	v = v & 128
	return v == 128
}
