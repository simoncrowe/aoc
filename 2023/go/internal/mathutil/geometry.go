package mathutil

type Vec2 struct {
	X, Y int
}

func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func (v Vec2) Flip() Vec2 {
	return Vec2{v.X - (2 * v.X), v.Y - (2 * v.Y)}
}

