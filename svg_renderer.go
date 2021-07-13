package main

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type Point struct {
	x int
	y int
}

type SVGRenderer struct {
	Value int
}

func (s SVGRenderer) Render() string {
	var sb strings.Builder

	width := 200
	height := 280
	padding := 10
	lineStyle := `stroke="#000" stroke-width="10"`
	polyLineStyle := fmt.Sprintf(`%s fill="none"`, lineStyle)

	ones := s.Value % 10
	tens := s.Value / 10 % 10
	hundreds := s.Value / 100 % 10
	thousands := s.Value / 1000 % 10

	canvas := svg.New(&sb)
	canvas.Start(width, height)

	// Pre-define anchor points
	p0 := Point{padding, padding}
	p1 := Point{width / 2, padding}
	p2 := Point{width - padding, padding}

	p3 := Point{padding, height / 3}
	p4 := Point{width / 2, height / 3}
	p5 := Point{width - padding, height / 3}

	p6 := Point{padding, height / 3 * 2}
	p7 := Point{width / 2, height / 3 * 2}
	p8 := Point{width - padding, height / 3 * 2}

	p9 := Point{padding, height - padding}
	p10 := Point{width / 2, height - padding}
	p11 := Point{width - padding, height - padding}

	xCoords := []int{p10.x, p1.x}
	yCoords := []int{p10.y, p1.y}

	switch ones {
	case 1:
		xCoords = append(xCoords, p2.x)
		yCoords = append(yCoords, p2.y)
	case 2:
		xCoords = append(xCoords, p4.x, p5.x)
		yCoords = append(yCoords, p4.y, p5.y)
	case 3:
		xCoords = append(xCoords, p5.x)
		yCoords = append(yCoords, p5.y)
	case 4:
		xCoords = append(xCoords, p4.x, p2.x)
		yCoords = append(yCoords, p4.y, p2.y)
	case 5:
		xCoords = append(xCoords, p2.x, p4.x)
		yCoords = append(yCoords, p2.y, p4.y)
	case 6:
		canvas.Line(p2.x, p2.y, p5.x, p5.y, lineStyle)
	case 7:
		xCoords = append(xCoords, p2.x, p5.x)
		yCoords = append(yCoords, p2.y, p5.y)
	case 8:
		xCoords = append(xCoords, p4.x, p5.x, p2.x)
		yCoords = append(yCoords, p4.y, p5.y, p2.y)
	case 9:
		xCoords = append(xCoords, p2.x, p5.x, p4.x)
		yCoords = append(yCoords, p2.y, p5.y, p4.y)
	}

	canvas.Polyline(xCoords, yCoords, polyLineStyle)

	xCoords = []int{p10.x, p1.x}
	yCoords = []int{p10.y, p1.y}

	switch tens {
	case 1:
		xCoords = append(xCoords, p0.x)
		yCoords = append(yCoords, p0.y)
	case 2:
		xCoords = append(xCoords, p4.x, p3.x)
		yCoords = append(yCoords, p4.y, p3.y)
	case 3:
		xCoords = append(xCoords, p3.x)
		yCoords = append(yCoords, p3.y)
	case 4:
		xCoords = append(xCoords, p4.x, p0.x)
		yCoords = append(yCoords, p4.y, p0.y)
	case 5:
		xCoords = append(xCoords, p0.x, p4.x)
		yCoords = append(yCoords, p0.y, p4.y)
	case 6:
		canvas.Line(p0.x, p0.y, p3.x, p3.y, lineStyle)
	case 7:
		xCoords = append(xCoords, p0.x, p3.x)
		yCoords = append(yCoords, p0.y, p3.y)
	case 8:
		xCoords = append(xCoords, p4.x, p3.x, p0.x)
		yCoords = append(yCoords, p4.y, p3.y, p0.y)
	case 9:
		xCoords = append(xCoords, p0.x, p3.x, p4.x)
		yCoords = append(yCoords, p0.y, p3.y, p4.y)
	}

	canvas.Polyline(xCoords, yCoords, polyLineStyle)

	xCoords = []int{p1.x, p10.x}
	yCoords = []int{p1.y, p10.y}

	switch hundreds {
	case 1:
		xCoords = append(xCoords, p11.x)
		yCoords = append(yCoords, p11.y)
	case 2:
		xCoords = append(xCoords, p7.x, p8.x)
		yCoords = append(yCoords, p7.y, p8.y)
	case 3:
		xCoords = append(xCoords, p8.x)
		yCoords = append(yCoords, p8.y)
	case 4:
		xCoords = append(xCoords, p7.x, p11.x)
		yCoords = append(yCoords, p7.y, p11.y)
	case 5:
		xCoords = append(xCoords, p11.x, p7.x)
		yCoords = append(yCoords, p11.y, p7.y)
	case 6:
		canvas.Line(p8.x, p8.y, p11.x, p11.y, lineStyle)
	case 7:
		xCoords = append(xCoords, p11.x, p8.x)
		yCoords = append(yCoords, p11.y, p8.y)
	case 8:
		xCoords = append(xCoords, p7.x, p8.x, p11.x)
		yCoords = append(yCoords, p7.y, p8.y, p11.y)
	case 9:
		xCoords = append(xCoords, p11.x, p8.x, p7.x)
		yCoords = append(yCoords, p11.y, p8.y, p7.y)
	}

	canvas.Polyline(xCoords, yCoords, polyLineStyle)

	xCoords = []int{p1.x, p10.x}
	yCoords = []int{p1.y, p10.y}

	switch thousands {
	case 1:
		xCoords = append(xCoords, p9.x)
		yCoords = append(yCoords, p9.y)
	case 2:
		xCoords = append(xCoords, p7.x, p6.x)
		yCoords = append(yCoords, p7.y, p6.y)
	case 3:
		xCoords = append(xCoords, p6.x)
		yCoords = append(yCoords, p6.y)
	case 4:
		xCoords = append(xCoords, p7.x, p9.x)
		yCoords = append(yCoords, p7.y, p9.y)
	case 5:
		xCoords = append(xCoords, p9.x, p7.x)
		yCoords = append(yCoords, p9.y, p7.y)
	case 6:
		canvas.Line(p6.x, p6.y, p9.x, p9.y, lineStyle)
	case 7:
		xCoords = append(xCoords, p9.x, p6.x)
		yCoords = append(yCoords, p9.y, p6.y)
	case 8:
		xCoords = append(xCoords, p7.x, p6.x, p9.x)
		yCoords = append(yCoords, p7.y, p6.y, p9.y)
	case 9:
		xCoords = append(xCoords, p9.x, p6.x, p7.x)
		yCoords = append(yCoords, p9.y, p6.y, p7.y)
	}

	canvas.Polyline(xCoords, yCoords, polyLineStyle)

	canvas.End()

	return sb.String()
}
