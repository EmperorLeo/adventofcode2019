package util

import (
	"fmt"
	"math"
)

type Coord struct {
	X, Y int
}

type Line struct {
	P1, P2 Coord
}

/*Intersects - this is only gonna work for horizontal and vertical lines*/
func (l Line) Intersects(l2 Line) bool {
	// LOL this code is so bad
	if l.isHorizontal() && !l2.isHorizontal() {
		if (l.P1.Y <= l2.P1.Y && l.P1.Y >= l2.P2.Y) || (l.P1.Y >= l2.P1.Y && l.P1.Y <= l2.P2.Y) {
			if (l2.P1.X <= l.P1.X && l2.P1.X >= l.P2.X) || (l2.P1.X >= l.P1.X && l2.P1.X <= l.P2.X) {
				return true
			}
		}
	}
	if !l.isHorizontal() && l2.isHorizontal() {
		if (l2.P1.Y <= l.P1.Y && l2.P1.Y >= l.P2.Y) || (l2.P1.Y >= l.P1.Y && l2.P1.Y <= l.P2.Y) {
			if (l.P1.X <= l2.P1.X && l.P1.X >= l2.P2.X) || (l.P1.X >= l2.P1.X && l.P1.X <= l2.P2.X) {
				return true
			}
		}
	}

	return false
}

func (l Line) Intersection(l2 Line) Coord {

	var coord Coord

	if l.isHorizontal() && !l2.isHorizontal() {
		if (l.P1.Y <= l2.P1.Y && l.P1.Y >= l2.P2.Y) || (l.P1.Y >= l2.P1.Y && l.P1.Y <= l2.P2.Y) {
			if (l2.P1.X <= l.P1.X && l2.P1.X >= l.P2.X) || (l2.P1.X >= l.P1.X && l2.P1.X <= l.P2.X) {
				coord = Coord{l2.P1.X, l.P1.Y}
			}
		}
	}
	if !l.isHorizontal() && l2.isHorizontal() {
		if (l2.P1.Y <= l.P1.Y && l2.P1.Y >= l.P2.Y) || (l2.P1.Y >= l.P1.Y && l2.P1.Y <= l.P2.Y) {
			if (l.P1.X <= l2.P1.X && l.P1.X >= l2.P2.X) || (l.P1.X >= l2.P1.X && l.P1.X <= l2.P2.X) {
				coord = Coord{l.P1.X, l2.P1.Y}
			}
		}
	}

	// fmt.Println("--- Intersection ---")

	// l.Print()
	// l2.Print()
	// coord.Print()

	return coord
}

func (c Coord) Manhattan(c2 Coord) int {
	ans := (int)(math.Abs((float64)(c.X-c2.X)) + math.Abs((float64)(c.Y-c2.Y)))
	// fmt.Printf("Manhattan Dist of {%d,%d} to {%d, %d}: %d\n", c.X, c.Y, c2.X, c2.Y, ans)
	return ans
}

func (l Line) isHorizontal() bool {
	// l.Print()
	return l.P1.Y == l.P2.Y
}

func (l Line) Length() int {
	return l.P1.Manhattan(l.P2)
}

func (c Coord) Print() {
	fmt.Printf("{%d,%d}\n", c.X, c.Y)
}

func (l Line) Print() {
	fmt.Printf("{%d,%d} -> {%d, %d} Is horizontal = %v\n", l.P1.X, l.P1.Y, l.P2.X, l.P2.Y, l.P1.Y == l.P2.Y)
}

func (c Coord) GetQuadrant() int {
	if c.X >= 0 && c.Y < 0 {
		return 1
	} else if c.X >= 0 && c.Y >= 0 {
		return 2
	} else if c.X < 0 && c.Y >= 0 {
		return 3
	} else {
		return 4
	}
}

func (c Coord) SlopeWith(o Coord) Coord {
	diffY := o.Y - c.Y
	diffX := o.X - c.X
	if diffY == 0 {
		return Coord{diffX / Abs(diffX), 0}
	} else if diffX == 0 {
		return Coord{0, diffY / Abs(diffY)}
	}
	reducer := gcf(Abs(diffX), Abs(diffY))
	diffY /= reducer
	diffX /= reducer

	return Coord{diffX, diffY}
}

func (c Coord) AbsoluteSlope() float64 {
	if c.X == 0 {
		return math.Inf(1)
	}
	return math.Abs(float64(c.Y) / float64(c.X))
}

func gcf(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
