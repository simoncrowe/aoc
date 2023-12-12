package mathutil

import (
	"testing"

	"github.com/EliCDavis/vector/vector2"
)

func TestIntersects(t *testing.T) {
	testCases := []struct {
		name            string
		segA            LineSeg
		segB            LineSeg
		shouldIntersect bool
	}{
		{"Parallel",
			LineSeg{vector2.New(1.0, 2.0), vector2.New(3.0, 2.0)},
			LineSeg{vector2.New(1.0, 1.0), vector2.New(3.0, 1.0)},
			false},
		{"Cross",
			LineSeg{vector2.New(10.0, 0.0), vector2.New(0.0, 10.0)},
			LineSeg{vector2.New(0.0, 0.0), vector2.New(10.0, 10.0)},
			true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Intersects(tc.segA, tc.segB)
			if result != tc.shouldIntersect {
				t.Errorf("%s: expected %v intersection with %v to be %t", tc.name, tc.segA, tc.segB, tc.shouldIntersect)
			}
		})
	}
}
