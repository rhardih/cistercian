package main

import "strings"

type TextGlyph struct {
	Ones      int
	Tens      int
	Hundreds  int
	Thousands int
}

func NewTextGlyph(value int) TextGlyph {
	ones := value % 10
	tens := value / 10 % 10
	hundreds := value / 100 % 10
	thousands := value / 1000 % 10

	return TextGlyph{ones, tens, hundreds, thousands}
}

func (a *TextGlyph) String() string {
	boxChars := []rune{
		'─', '│', '┌', '┐', '└', '┘',
		'├', '┤', '┬', '┴', '┼', '╱',
		'╲', '╴', '╵', '╶', '╷',
	}

	rows := 13
	cols := 9
	glyph := make([]rune, rows*cols)

	// Add center staves, which are always there
	for row := 0; row < rows; row += 1 {
		char := boxChars[1]

		if row == 0 {
			char = boxChars[16]
		}

		if row == 12 {
			char = boxChars[14]
		}

		for col := 0; col < cols; col += 1 {
			if col == 4 {
				glyph[row*cols+col] = char
			} else {
				glyph[row*cols+col] = ' '
			}
		}
	}

	switch a.Ones {
	case 1:
		offset := 4

		glyph[offset] = boxChars[2]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[0]
	case 2:
		offset := 4 + cols*4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[0]
	case 3:
		offset := 5

		glyph[offset+cols] = boxChars[12]
		glyph[offset+cols*2+1] = boxChars[12]
		glyph[offset+cols*3+2] = boxChars[12]
	case 4:
		offset := 8

		glyph[offset+cols-1] = boxChars[11]
		glyph[offset+cols*2-2] = boxChars[11]
		glyph[offset+cols*3-3] = boxChars[11]
	case 5:
		offset := 4

		glyph[offset] = boxChars[2]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[0]

		offset = 8
		glyph[offset+cols-1] = boxChars[11]
		glyph[offset+cols*2-2] = boxChars[11]
		glyph[offset+cols*3-3] = boxChars[11]
	case 6:
		offset := 8

		glyph[offset] = boxChars[16]
		glyph[offset+cols] = boxChars[1]
		glyph[offset+cols*2] = boxChars[1]
		glyph[offset+cols*3] = boxChars[1]
		glyph[offset+cols*4] = boxChars[14]
	case 7:
		offset := 4

		glyph[offset] = boxChars[2]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[3]

		offset = 8

		glyph[offset+cols] = boxChars[1]
		glyph[offset+cols*2] = boxChars[1]
		glyph[offset+cols*3] = boxChars[1]
		glyph[offset+cols*4] = boxChars[14]
	case 8:
		offset := 8

		glyph[offset] = boxChars[16]
		glyph[offset+cols] = boxChars[1]
		glyph[offset+cols*2] = boxChars[1]
		glyph[offset+cols*3] = boxChars[1]

		offset = 4 + cols*4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[5]
	case 9:
		offset := 4

		glyph[offset] = boxChars[2]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[3]

		offset = 8

		glyph[offset+cols] = boxChars[1]
		glyph[offset+cols*2] = boxChars[1]
		glyph[offset+cols*3] = boxChars[1]

		offset = 4 + cols*4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[5]
	default:
		// zero, do nothing
	}

	switch a.Tens {
	case 1:
		glyph[0] = boxChars[0]
		glyph[1] = boxChars[0]
		glyph[2] = boxChars[0]
		glyph[3] = boxChars[0]

		if glyph[4] == boxChars[16] {
			glyph[4] = boxChars[3]
		} else {
			glyph[4] = boxChars[8]
		}
	case 2:
		offset := cols * 4

		glyph[offset] = boxChars[0]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]

		if glyph[offset+4] == boxChars[6] {
			glyph[offset+4] = boxChars[10]
		} else {
			glyph[offset+4] = boxChars[7]
		}
	case 3:
		offset := 3

		glyph[offset+cols] = boxChars[11]
		glyph[offset+cols*2-1] = boxChars[11]
		glyph[offset+cols*3-2] = boxChars[11]
	case 4:
		glyph[cols+1] = boxChars[12]
		glyph[cols*2+2] = boxChars[12]
		glyph[cols*3+3] = boxChars[12]
	case 5:
		glyph[0] = boxChars[0]
		glyph[1] = boxChars[0]
		glyph[2] = boxChars[0]
		glyph[3] = boxChars[0]

		if glyph[4] == boxChars[16] {
			glyph[4] = boxChars[3]
		} else {
			glyph[4] = boxChars[8]
		}

		glyph[cols+1] = boxChars[12]
		glyph[cols*2+2] = boxChars[12]
		glyph[cols*3+3] = boxChars[12]
	case 6:
		glyph[0] = boxChars[16]
		glyph[cols] = boxChars[1]
		glyph[cols*2] = boxChars[1]
		glyph[cols*3] = boxChars[1]
		glyph[cols*4] = boxChars[14]
	case 7:
		glyph[0] = boxChars[2]
		glyph[1] = boxChars[0]
		glyph[2] = boxChars[0]
		glyph[3] = boxChars[0]

		if glyph[4] == boxChars[16] {
			glyph[4] = boxChars[3]
		} else {
			glyph[4] = boxChars[8]
		}

		glyph[cols] = boxChars[1]
		glyph[cols*2] = boxChars[1]
		glyph[cols*3] = boxChars[1]
		glyph[cols*4] = boxChars[14]
	case 8:
		glyph[0] = boxChars[16]
		glyph[cols] = boxChars[1]
		glyph[cols*2] = boxChars[1]
		glyph[cols*3] = boxChars[1]

		offset := cols * 4

		glyph[offset] = boxChars[4]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]

		if glyph[offset+4] == boxChars[6] {
			glyph[offset+4] = boxChars[10]
		} else {
			glyph[offset+4] = boxChars[7]
		}
	case 9:
		glyph[0] = boxChars[2]
		glyph[1] = boxChars[0]
		glyph[2] = boxChars[0]
		glyph[3] = boxChars[0]

		if glyph[4] == boxChars[16] {
			glyph[4] = boxChars[3]
		} else {
			glyph[4] = boxChars[8]
		}

		glyph[cols] = boxChars[1]
		glyph[cols*2] = boxChars[1]
		glyph[cols*3] = boxChars[1]

		offset := cols * 4

		glyph[offset] = boxChars[4]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]

		if glyph[offset+4] == boxChars[6] {
			glyph[offset+4] = boxChars[10]
		} else {
			glyph[offset+4] = boxChars[7]
		}
	default:
		// zero, do nothing
	}

	switch a.Hundreds {
	case 1:
		offset := cols*(rows-1) + 4

		glyph[offset] = boxChars[4]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[0]
	case 2:
		offset := cols*(rows-5) + 4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[0]
	case 3:
		offset := cols*(rows-2) + 5

		glyph[offset] = boxChars[11]
		glyph[offset-cols+1] = boxChars[11]
		glyph[offset-cols*2+2] = boxChars[11]
	case 4:
		offset := cols*(rows-4) + 5

		glyph[offset] = boxChars[12]
		glyph[offset+cols+1] = boxChars[12]
		glyph[offset+cols*2+2] = boxChars[12]
	case 5:
		offset := cols*(rows-4) + 5

		glyph[offset] = boxChars[12]
		glyph[offset+cols+1] = boxChars[12]
		glyph[offset+cols*2+2] = boxChars[12]

		offset = cols*(rows-1) + 4

		glyph[offset] = boxChars[4]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[0]
	case 6:
		offset := rows*cols - 1

		glyph[offset] = boxChars[14]
		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]
		glyph[offset-cols*4] = boxChars[16]
	case 7:
		offset := rows*cols - 1

		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]
		glyph[offset-cols*4] = boxChars[16]

		offset = cols*(rows-1) + 4

		glyph[offset] = boxChars[4]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[5]
	case 8:
		offset := cols*(rows-5) + 4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[3]

		offset = rows*cols - 1

		glyph[offset] = boxChars[14]
		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]
	case 9:
		offset := cols*(rows-5) + 4

		glyph[offset] = boxChars[6]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[3]

		offset = rows*cols - 1

		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]

		offset = cols*(rows-1) + 4

		glyph[offset] = boxChars[4]
		glyph[offset+1] = boxChars[0]
		glyph[offset+2] = boxChars[0]
		glyph[offset+3] = boxChars[0]
		glyph[offset+4] = boxChars[5]
	default:
		// zero, do nothing
	}

	switch a.Thousands {
	case 1:
		offset := cols*rows - 5

    if glyph[offset] == boxChars[4] {
      glyph[offset] = boxChars[9]
    } else {
      glyph[offset] = boxChars[5]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[0]
	case 2:
		offset := cols*(rows-4) - 5

    if glyph[offset] == boxChars[6] {
      glyph[offset] = boxChars[10]
    } else {
      glyph[offset] = boxChars[7]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[0]
	case 3:
		offset := cols*(rows-1) - 6

		glyph[offset] = boxChars[12]
		glyph[offset-cols-1] = boxChars[12]
		glyph[offset-cols*2-2] = boxChars[12]
	case 4:
		offset := cols*(rows-1) - 8

		glyph[offset] = boxChars[11]
		glyph[offset-cols+1] = boxChars[11]
		glyph[offset-cols*2+2] = boxChars[11]
	case 5:
		offset := cols*(rows-1) - 8

		glyph[offset] = boxChars[11]
		glyph[offset-cols+1] = boxChars[11]
		glyph[offset-cols*2+2] = boxChars[11]

		offset = cols*rows - 5

    if glyph[offset] == boxChars[4] {
      glyph[offset] = boxChars[9]
    } else {
      glyph[offset] = boxChars[5]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[0]
	case 6:
		offset := cols * (rows - 1)

		glyph[offset] = boxChars[14]
		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]
		glyph[offset-cols*4] = boxChars[16]
	case 7:
		offset := cols*rows - 5

    if glyph[offset] == boxChars[4] {
      glyph[offset] = boxChars[9]
    } else {
      glyph[offset] = boxChars[5]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[4]

		offset = cols * (rows - 1)

		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]
		glyph[offset-cols*4] = boxChars[16]
	case 8:
		offset := cols*(rows-4) - 5

    if glyph[offset] == boxChars[6] {
      glyph[offset] = boxChars[10]
    } else {
      glyph[offset] = boxChars[7]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[2]

		offset = cols * (rows - 1)

		glyph[offset] = boxChars[14]
		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]
	case 9:
		offset := cols*(rows-4) - 5

    if glyph[offset] == boxChars[6] {
      glyph[offset] = boxChars[10]
    } else {
      glyph[offset] = boxChars[7]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[2]

		offset = cols * (rows - 1)

		glyph[offset-cols] = boxChars[1]
		glyph[offset-cols*2] = boxChars[1]
		glyph[offset-cols*3] = boxChars[1]

		offset = cols*rows - 5

    if glyph[offset] == boxChars[4] {
      glyph[offset] = boxChars[9]
    } else {
      glyph[offset] = boxChars[5]
    }

		glyph[offset-1] = boxChars[0]
		glyph[offset-2] = boxChars[0]
		glyph[offset-3] = boxChars[0]
		glyph[offset-4] = boxChars[4]
	default:
		// zero, do nothing
	}

	sb := strings.Builder{}

	for row := 0; row < rows; row += 1 {
		if row > 0 {
			sb.WriteString("\n")
		}

		for col := 0; col < cols; col += 1 {
			sb.WriteRune(glyph[row*cols+col])
		}
	}

	return sb.String()
}
