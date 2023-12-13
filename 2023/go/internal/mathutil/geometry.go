package mathutil

import (
	"github.com/EliCDavis/vector/vector2"
)

type LineSeg struct {
	P, Q vector2.Vector[float64]
}

func CountIntersections(p LineSeg, contour []LineSeg) int {
	count := 0
	for _, q := range contour {
		if Intersects(p, q) {
			count++
		}
	}
	return count

}
func Intersects(one, two LineSeg) bool {
	o1 := getOrientation(one.P, one.Q, two.P)
	o2 := getOrientation(one.P, one.Q, two.Q)
	o3 := getOrientation(two.P, two.Q, one.P)
	o4 := getOrientation(two.P, two.Q, one.Q)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == 0 && onSegment(one.P, two.P, one.Q) {
		return true
	}
	if o2 == 0 && onSegment(one.P, two.Q, one.Q) {
		return true
	}
	if o3 == 0 && onSegment(two.P, one.P, two.Q) {
		return true
	}
	if o4 == 0 && onSegment(two.P, one.Q, two.Q) {
		return true
	}

	return false
}

// Find orentation of ordered triplet p,q,r
func getOrientation(p, q, r vector2.Vector[float64]) float64 {
	val := (q.Y()-p.Y())*(r.X()-q.X()) - (q.X()-p.X())*(r.Y()-q.Y())
	if val > 0 {
		// Clockwise
		return 1
	} else if val < 0 {
		// Counterclockwise
		return -1 
	} else {
		// Colinear
		return 0
	}
}

// p, q, r are colinear; does q lie on pr?
func onSegment(p, q, r vector2.Vector[float64]) bool {
	onSegX := q.X() <= max(p.X(), r.X()) && q.X() >= min(p.X(), r.X())
	onSegY := q.Y() <= max(p.Y(), r.Y()) && q.Y() >= min(p.Y(), r.Y())
	return onSegX && onSegY
}
