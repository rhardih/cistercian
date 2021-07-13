package main

import (
	"strings"
	"testing"
)

func TestTextRenderer(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		renderer := TextRenderer{1}

		actual := renderer.Render()

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
		renderer := TextRenderer{2}

		actual := renderer.Render()

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
		renderer := TextRenderer{3}

		actual := renderer.Render()

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
		renderer := TextRenderer{4}

		actual := renderer.Render()

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
		renderer := TextRenderer{5}

		actual := renderer.Render()

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
		renderer := TextRenderer{6}

		actual := renderer.Render()

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
		renderer := TextRenderer{7}

		actual := renderer.Render()

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
		renderer := TextRenderer{8}

		actual := renderer.Render()

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
		renderer := TextRenderer{9}

		actual := renderer.Render()

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
		renderer := TextRenderer{10}

		actual := renderer.Render()

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
		renderer := TextRenderer{20}

		actual := renderer.Render()

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
		renderer := TextRenderer{30}

		actual := renderer.Render()

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
		renderer := TextRenderer{40}

		actual := renderer.Render()

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
		renderer := TextRenderer{50}

		actual := renderer.Render()

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
		renderer := TextRenderer{60}

		actual := renderer.Render()

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
		renderer := TextRenderer{70}

		actual := renderer.Render()

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
		renderer := TextRenderer{80}

		actual := renderer.Render()

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
		renderer := TextRenderer{90}

		actual := renderer.Render()

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
		renderer := TextRenderer{100}

		actual := renderer.Render()

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
		renderer := TextRenderer{200}

		actual := renderer.Render()

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
		renderer := TextRenderer{300}

		actual := renderer.Render()

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
		renderer := TextRenderer{400}

		actual := renderer.Render()

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
		renderer := TextRenderer{500}

		actual := renderer.Render()

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
		renderer := TextRenderer{600}

		actual := renderer.Render()

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
		renderer := TextRenderer{700}

		actual := renderer.Render()

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
		renderer := TextRenderer{800}

		actual := renderer.Render()

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
		renderer := TextRenderer{900}

		actual := renderer.Render()

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
		renderer := TextRenderer{1000}

		actual := renderer.Render()

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
		renderer := TextRenderer{2000}

		actual := renderer.Render()

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
		renderer := TextRenderer{3000}

		actual := renderer.Render()

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
		renderer := TextRenderer{4000}

		actual := renderer.Render()

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
		renderer := TextRenderer{5000}

		actual := renderer.Render()

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
		renderer := TextRenderer{6000}

		actual := renderer.Render()

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
		renderer := TextRenderer{7000}

		actual := renderer.Render()

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
		renderer := TextRenderer{8000}

		actual := renderer.Render()

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
		renderer := TextRenderer{9000}

		actual := renderer.Render()

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
			renderer := TextRenderer{1111}

			actual := renderer.Render()

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
			renderer := TextRenderer{2222}

			actual := renderer.Render()

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
			renderer := TextRenderer{5555}

			actual := renderer.Render()

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
			renderer := TextRenderer{7777}

			actual := renderer.Render()

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
			renderer := TextRenderer{8888}

			actual := renderer.Render()

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
			renderer := TextRenderer{9999}

			actual := renderer.Render()

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
