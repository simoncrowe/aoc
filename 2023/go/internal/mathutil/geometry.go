package mathutil

import (
	"github.com/EliCDavis/vector/vector2"
)


type LineSeg struct {
	A, B vector2.Vector[float32]
}


func Intersects(p, q LineSeg) bool {
	o1 := getOrientation(p.A, q.A, p.B)
	o2 := getOrientation(p.A, q.A, q.B)
	o3 := getOrientation(p.B, q.B, p.A)
	o4 := getOrientation(p.B, q.B, q.A)

	if o1 != o2 && o3 != o4 {
		return true
	}

	if o1 == 0 && onSegment(p.A, p.B, q.A) {
		return true
	}
	if o2 == 0 && onSegment(p.A, q.B, q.A) {
		return true
	}
	if o3 == 0 && onSegment(p.B, p.A, q.B) {
		return true
	}
	if o4 == 0 && onSegment(p.B, q.A, q.B) {
		return true
	}
	
	return false
}

// Find orentation of ordered triplet p,q,r
func getOrientation(p, q, r vector2.Vector[float32]) float32 {
	val := ((q.Y() - p.Y()) * (r.X() - q.X())) - ((q.X()-p.X()) * (r.Y()-q.Y()))

	if val > 0 {
		// Clockwise
		return 1
	} else if val < 0 {
		// Counterclockwise
		return 2
	} else {
		// Colomera
		return 0
	}
}

// p, q, r are colinear; does q lie on pr?
func onSegment(p, q, r vector2.Vector[float32]) bool {
	onSegX := q.X() <= max(p.X(), r.X()) && q.X() >= min(p.X(), r.X()) 
	onSegY := q.Y() <= min(p.Y(), r.Y()) && q.Y() >= min(p.Y(), r.Y())
	return onSegX && onSegY
}

