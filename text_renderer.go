package main

import "fmt"

type TextRenderer struct {
	Value int
}

func (a TextRenderer) Output() {
	glyph := NewTextGlyph(a.Value)

	fmt.Println(glyph.String())
}
