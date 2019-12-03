package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

type runningDist struct {
	l util.Line
	d int
}

func Silver() {
	input := util.ReadLines(3)
	strWire1 := input[0]
	strWire2 := input[1]
	directions1 := strings.Split(strWire1, ",")
	directions2 := strings.Split(strWire2, ",")
	origin := util.Coord{0, 0}

	wire1 := make([]runningDist, len(directions1))
	curCoord := origin
	curSteps := 0
	for i, dir1 := range directions1 {
		newLine := parse(curCoord, dir1)
		curSteps += newLine.Length()
		rDist := runningDist{newLine, curSteps}
		wire1[i] = rDist
		curCoord = newLine.P2
	}

	wire2 := make([]runningDist, len(directions2))
	curCoord = origin
	curSteps = 0
	for i, dir2 := range directions2 {
		newLine := parse(curCoord, dir2)
		curSteps += newLine.Length()
		rDist := runningDist{newLine, curSteps}
		wire2[i] = rDist
		curCoord = newLine.P2
	}

	// initialize Silver Answer
	minDist := math.MaxInt32
	var best util.Coord

	// initialize Gold Answer
	lowestSteps := math.MaxInt32
	var lowestStepsCoord util.Coord

	// Loop through all potential intersections
	for _, w := range wire1 {
		wLine := w.l
		wRunning := w.d
		for _, x := range wire2 {
			xLine := x.l
			xRunning := x.d
			if wLine.Intersects(xLine) {
				// Find the intersection and append it
				intersection := wLine.Intersection(xLine)

				// If this is the closest intersection so far to the origin, update the best coord
				dist := origin.Manhattan(intersection)
				if dist < minDist && dist != 0 {
					minDist = dist
					best = intersection
				}

				// Attempt to get the real running distance by subtracting the stuff after intersection
				wRunning -= (wLine.Length() - wLine.P1.Manhattan(intersection))
				xRunning -= (xLine.Length() - xLine.P1.Manhattan(intersection))
				// replace lowest steps coord if applicable
				if wRunning+xRunning < lowestSteps && intersection != origin {
					lowestStepsCoord = intersection
					lowestSteps = wRunning + xRunning
				}
			}
		}
	}

	fmt.Printf("Best intersection is {%d, %d} with distance of %d\n", best.X, best.Y, origin.Manhattan(best))
	fmt.Printf("Least steps intersection is {%d, %d} with distance of %d\n", lowestStepsCoord.X, lowestStepsCoord.Y, lowestSteps)
}

func Gold() {
	/* Included in Silver */
}

func parse(c util.Coord, direction string) util.Line {
	dir := direction[0]
	remaining, _ := strconv.Atoi(direction[1:])
	var coord util.Coord
	switch dir {
	case 'R':
		coord = util.Coord{c.X + remaining, c.Y}
	case 'L':
		coord = util.Coord{c.X - remaining, c.Y}
	case 'D':
		coord = util.Coord{c.X, c.Y - remaining}
	case 'U':
		coord = util.Coord{c.X, c.Y + remaining}
	default:
		fmt.Println("Should not happen")
	}
	return util.Line{c, coord}
}
