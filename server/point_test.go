package server

import (
	"fmt"
	"math"
	"testing"
)

func Test_Point_Distance(t *testing.T) {
	for _, tc := range []struct {
		p1   *Point
		p2   *Point
		dist float32
	}{{
		p1:   &Point{X: 0, Y: 0},
		p2:   &Point{X: 2, Y: 0},
		dist: 2,
	}, {
		p1:   &Point{X: 0, Y: 0},
		p2:   &Point{X: 3, Y: 4},
		dist: 5,
	}} {
		t.Run(fmt.Sprintf("%v_%v_%v", tc.p1, tc.p2, tc.dist), func(t *testing.T) {
			if d := tc.p1.Distance(tc.p2); !floatCompare(d, tc.dist, 0.001) {
				t.Errorf("wrong distance, want: %v got: %v", tc.dist, d)
			}
		})
	}
}

func Test_Point_CopyWithSameAlpha(t *testing.T) {
	for _, tc := range []struct {
		p1 *Point
		p2 *Point
		l  float32
	}{{
		p1: &Point{X: 3, Y: 0},
		p2: &Point{X: 5, Y: 0},
		l:  5,
	}, {
		p1: &Point{X: -2, Y: -2},
		p2: &Point{X: -1, Y: -1},
		l:  float32(math.Pow(2, 0.5)),
	}} {
		t.Run(fmt.Sprintf("%v_%v_%v", tc.p1, tc.p2, tc.l), func(t *testing.T) {
			if p := tc.p1.CopyWithSameAlpha(tc.l); !floatCompare(p.Distance(tc.p2), 0, 0.001) {
				t.Errorf("wrong result point, want: %v got: %v", tc.p2, p)
			}
		})
	}
}
