package overlapping_rectangles

import (
	"math"
)

type Point struct {
	X, Y float64
}

type Rectangle struct {
	BL, TR Point
}

func (r Rectangle) Area() float64 {
	d1 := r.TR.X - r.BL.X
	d2 := r.TR.Y - r.BL.Y
	return d1 * d2
}

func (r Rectangle) Intersect(b Rectangle) Rectangle {
	blx := math.Max(r.BL.X, b.BL.X)
	bly := math.Max(r.BL.Y, b.BL.Y)
	tlx := math.Min(r.TR.X, b.TR.X)
	tly := math.Min(r.TR.Y, b.TR.Y)
	if blx >= tlx || bly >= tly {
		return Rectangle{}
	}
	return Rectangle{Point{blx, bly}, Point{tlx, tly}}
}
