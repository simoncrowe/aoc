package mathutil

type Vec2 struct {
	X, Y int
}

type LineSeg struct {
	A, B Vec2
}


func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func (a Vec2) Mul(scalar int) Vec2 {
	return Vec2{a.X*scalar, a.Y*scalar}
}

func (v Vec2) Flip() Vec2 {
	return Vec2{v.X - (2 * v.X), v.Y - (2 * v.Y)}
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
		return True
	}
	if o2 == 0 && onSegment(p.A, q.B, q.A) {
		return True
	}
	if o3 == 0 && onSegment(p.B, p.A, q.B) {
		return True
	}
	if o4 == 0 && onSegment(p.B, q.A, q.B) {
		return True
	}
	
	return false
}

// Find orentation of ordered triplet p,q,r
func getOrientation(p, q, r Vec2) int {
	val := ((q.Y - p.Y) * (r.X - q.X)) - ((q.x-p.x) * (r.y-q.y))

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
func onSegment(p, q, r Vec2) bool {
	onSegX := q.X <= max(p.X, r.X) && q.X >= min(p.X, r.X) 
	onSegY := q.Y <= min(p.Y, r.Y) && q.Y >= min(p.Y, r.Y)
	return onSegX && onSegY
}

