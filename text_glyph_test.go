package main

import (
	"strings"
	"testing"
)

func TestTextGlyph(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		glyph := NewTextGlyph(1)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ┌────",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Two", func(t *testing.T) {
		glyph := NewTextGlyph(2)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ├────",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Three", func(t *testing.T) {
		glyph := NewTextGlyph(3)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │╲   ",
			"    │ ╲  ",
			"    │  ╲ ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Four", func(t *testing.T) {
		glyph := NewTextGlyph(4)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │  ╱ ",
			"    │ ╱  ",
			"    │╱   ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Five", func(t *testing.T) {
		glyph := NewTextGlyph(5)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ┌────",
			"    │  ╱ ",
			"    │ ╱  ",
			"    │╱   ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Six", func(t *testing.T) {
		glyph := NewTextGlyph(6)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷   ╷",
			"    │   │",
			"    │   │",
			"    │   │",
			"    │   ╵",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Seven", func(t *testing.T) {
		glyph := NewTextGlyph(7)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ┌───┐",
			"    │   │",
			"    │   │",
			"    │   │",
			"    │   ╵",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Eight", func(t *testing.T) {
		glyph := NewTextGlyph(8)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷   ╷",
			"    │   │",
			"    │   │",
			"    │   │",
			"    ├───┘",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Nine", func(t *testing.T) {
		glyph := NewTextGlyph(9)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ┌───┐",
			"    │   │",
			"    │   │",
			"    │   │",
			"    ├───┘",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Ten", func(t *testing.T) {
		glyph := NewTextGlyph(10)

		actual := glyph.String()

		expected := strings.Join([]string{
			"────┐    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Twenty", func(t *testing.T) {
		glyph := NewTextGlyph(20)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"────┤    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Thirty", func(t *testing.T) {
		glyph := NewTextGlyph(30)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"   ╱│    ",
			"  ╱ │    ",
			" ╱  │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Fourty", func(t *testing.T) {
		glyph := NewTextGlyph(40)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			" ╲  │    ",
			"  ╲ │    ",
			"   ╲│    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Fifty", func(t *testing.T) {
		glyph := NewTextGlyph(50)

		actual := glyph.String()

		expected := strings.Join([]string{
			"────┐    ",
			" ╲  │    ",
			"  ╲ │    ",
			"   ╲│    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Sixty", func(t *testing.T) {
		glyph := NewTextGlyph(60)

		actual := glyph.String()

		expected := strings.Join([]string{
			"╷   ╷    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"╵   │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Seventy", func(t *testing.T) {
		glyph := NewTextGlyph(70)

		actual := glyph.String()

		expected := strings.Join([]string{
			"┌───┐    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"╵   │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Eighty", func(t *testing.T) {
		glyph := NewTextGlyph(80)

		actual := glyph.String()

		expected := strings.Join([]string{
			"╷   ╷    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"└───┤    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Ninety", func(t *testing.T) {
		glyph := NewTextGlyph(90)

		actual := glyph.String()

		expected := strings.Join([]string{
			"┌───┐    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"└───┤    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("One hundred", func(t *testing.T) {
		glyph := NewTextGlyph(100)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    └────",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Two hundred", func(t *testing.T) {
		glyph := NewTextGlyph(200)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ├────",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Three hundred", func(t *testing.T) {
		glyph := NewTextGlyph(300)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │  ╱ ",
			"    │ ╱  ",
			"    │╱   ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Four hundred", func(t *testing.T) {
		glyph := NewTextGlyph(400)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │╲   ",
			"    │ ╲  ",
			"    │  ╲ ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Five hundred", func(t *testing.T) {
		glyph := NewTextGlyph(500)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │╲   ",
			"    │ ╲  ",
			"    │  ╲ ",
			"    └────",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Six hundred", func(t *testing.T) {
		glyph := NewTextGlyph(600)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │   ╷",
			"    │   │",
			"    │   │",
			"    │   │",
			"    ╵   ╵",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Seven hundred", func(t *testing.T) {
		glyph := NewTextGlyph(700)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │   ╷",
			"    │   │",
			"    │   │",
			"    │   │",
			"    └───┘",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Eight hundred", func(t *testing.T) {
		glyph := NewTextGlyph(800)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ├───┐",
			"    │   │",
			"    │   │",
			"    │   │",
			"    ╵   ╵",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Nine hundred", func(t *testing.T) {
		glyph := NewTextGlyph(900)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ├───┐",
			"    │   │",
			"    │   │",
			"    │   │",
			"    └───┘",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("One thousand", func(t *testing.T) {
		glyph := NewTextGlyph(1000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"────┘    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Two thousand", func(t *testing.T) {
		glyph := NewTextGlyph(2000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"────┤    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Three thousand", func(t *testing.T) {
		glyph := NewTextGlyph(3000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			" ╲  │    ",
			"  ╲ │    ",
			"   ╲│    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Four thousand", func(t *testing.T) {
		glyph := NewTextGlyph(4000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"   ╱│    ",
			"  ╱ │    ",
			" ╱  │    ",
			"    ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Five thousand", func(t *testing.T) {
		glyph := NewTextGlyph(5000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"   ╱│    ",
			"  ╱ │    ",
			" ╱  │    ",
			"────┘    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Six thousand", func(t *testing.T) {
		glyph := NewTextGlyph(6000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"╷   │    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"╵   ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Seven thousand", func(t *testing.T) {
		glyph := NewTextGlyph(7000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"╷   │    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"└───┘    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Eight thousand", func(t *testing.T) {
		glyph := NewTextGlyph(8000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"┌───┤    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"╵   ╵    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Nine thousand", func(t *testing.T) {
		glyph := NewTextGlyph(9000)

		actual := glyph.String()

		expected := strings.Join([]string{
			"    ╷    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"    │    ",
			"┌───┤    ",
			"│   │    ",
			"│   │    ",
			"│   │    ",
			"└───┘    ",
		}, "\n")

		if strings.Compare(actual, expected) != 0 {
			t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
		}
	})

	t.Run("Crosses", func(t *testing.T) {
		t.Run("Ones", func(t *testing.T) {
			glyph := NewTextGlyph(1111)

			actual := glyph.String()

			expected := strings.Join([]string{
				"────┬────",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"────┴────",
			}, "\n")

			if strings.Compare(actual, expected) != 0 {
				t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
			}
		})

		t.Run("Twos", func(t *testing.T) {
			glyph := NewTextGlyph(2222)

			actual := glyph.String()

			expected := strings.Join([]string{
				"    ╷    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"────┼────",
				"    │    ",
				"    │    ",
				"    │    ",
				"────┼────",
				"    │    ",
				"    │    ",
				"    │    ",
				"    ╵    ",
			}, "\n")

			if strings.Compare(actual, expected) != 0 {
				t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
			}
		})

		t.Run("Fives", func(t *testing.T) {
			glyph := NewTextGlyph(5555)

			actual := glyph.String()

			expected := strings.Join([]string{
				"────┬────",
				" ╲  │  ╱ ",
				"  ╲ │ ╱  ",
				"   ╲│╱   ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"    │    ",
				"   ╱│╲   ",
				"  ╱ │ ╲  ",
				" ╱  │  ╲ ",
				"────┴────",
			}, "\n")

			if strings.Compare(actual, expected) != 0 {
				t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
			}
		})

		t.Run("Sevens", func(t *testing.T) {
			glyph := NewTextGlyph(7777)

			actual := glyph.String()

			expected := strings.Join([]string{
				"┌───┬───┐",
				"│   │   │",
				"│   │   │",
				"│   │   │",
				"╵   │   ╵",
				"    │    ",
				"    │    ",
				"    │    ",
				"╷   │   ╷",
				"│   │   │",
				"│   │   │",
				"│   │   │",
				"└───┴───┘",
			}, "\n")

			if strings.Compare(actual, expected) != 0 {
				t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
			}
		})

		t.Run("Eights", func(t *testing.T) {
			glyph := NewTextGlyph(8888)

			actual := glyph.String()

			expected := strings.Join([]string{
				"╷   ╷   ╷",
				"│   │   │",
				"│   │   │",
				"│   │   │",
				"└───┼───┘",
				"    │    ",
				"    │    ",
				"    │    ",
				"┌───┼───┐",
				"│   │   │",
				"│   │   │",
				"│   │   │",
				"╵   ╵   ╵",
			}, "\n")

			if strings.Compare(actual, expected) != 0 {
				t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
			}
		})

		t.Run("Nines", func(t *testing.T) {
			glyph := NewTextGlyph(9999)

			actual := glyph.String()

			expected := strings.Join([]string{
				"┌───┬───┐",
				"│   │   │",
				"│   │   │",
				"│   │   │",
				"└───┼───┘",
				"    │    ",
				"    │    ",
				"    │    ",
				"┌───┼───┐",
				"│   │   │",
				"│   │   │",
				"│   │   │",
				"└───┴───┘",
			}, "\n")

			if strings.Compare(actual, expected) != 0 {
				t.Errorf("\nActual:\n'%s'\n !== \nExpected:\n'%s'", actual, expected)
			}
		})
	})
}
