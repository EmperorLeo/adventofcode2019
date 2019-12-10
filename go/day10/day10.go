package day10

import (
	"fmt"
	"sort"

	"github.com/EmperorLeo/adventofcode2019/util"
)

const debug bool = false

type vectorLoc struct {
	location, vector util.Coord
}

type asteroidNode struct {
	location, vector util.Coord
	prev, next       *asteroidNode
}

func Silver() {
	input := util.ReadLines(10)

	ans, count, _ := getBestStationPosition(input)

	fmt.Printf("The best asteroid for the station is at (%d, %d) with %d asteroids in sight\n", ans.X, ans.Y, count)
}

func Gold() {
	input := util.ReadLines(10)

	station, _, asteroids := getBestStationPosition(input)

	result := getNthAsteroid(station, asteroids, 200)

	fmt.Printf("200 asteroid value (%d, %d) = %d\n", result.X, result.Y, result.X*100+result.Y)
}

func getBestStationPosition(field []string) (util.Coord, int, []util.Coord) {
	asteroidList := []util.Coord{}

	for y, r := range field {
		for x, spot := range r {
			if spot == '#' {
				asteroidList = append(asteroidList, util.Coord{x, y})
			}
		}
	}

	var bestAsteroid util.Coord
	var bestAsteroidCount int

	for i := 0; i < len(asteroidList); i++ {
		a1 := asteroidList[i]
		var uniqueSlopes int
		slopes := map[util.Coord]int{}

		for j := 0; j < len(asteroidList); j++ {
			a2 := asteroidList[j]
			if a1 != a2 {
				slope := a1.SlopeWith(a2)
				if slopes[slope] == 0 {
					uniqueSlopes++
				}
				slopes[slope]++
			}
		}

		if uniqueSlopes > bestAsteroidCount {
			bestAsteroidCount = uniqueSlopes
			bestAsteroid = a1
		}
	}

	return bestAsteroid, bestAsteroidCount, asteroidList
}

func getNthAsteroid(station util.Coord, asteroids []util.Coord, n int) util.Coord {
	asteroidVectors := []vectorLoc{}
	for _, asteroid := range asteroids {
		if station != asteroid {
			vector := station.SlopeWith(asteroid)
			asteroidVectors = append(asteroidVectors, vectorLoc{asteroid, vector})
		}
	}
	sort.Slice(asteroidVectors, func(i, j int) bool {
		a := asteroidVectors[i]
		b := asteroidVectors[j]
		aQuad := a.vector.GetQuadrant()
		bQuad := b.vector.GetQuadrant()
		if aQuad != bQuad {
			return aQuad < bQuad
		}

		if a.vector == b.vector {
			return a.location.Manhattan(station) < b.location.Manhattan(station)
		}

		switch aQuad {
		case 1:
			return a.vector.AbsoluteSlope() > b.vector.AbsoluteSlope()
		case 2:
			return a.vector.AbsoluteSlope() < b.vector.AbsoluteSlope()
		case 3:
			return a.vector.AbsoluteSlope() > b.vector.AbsoluteSlope()
		case 4:
			return a.vector.AbsoluteSlope() < b.vector.AbsoluteSlope()
		default:
			return true
		}
	})

	if debug {
		for _, a := range asteroidVectors {
			fmt.Printf("Asteroid (%d, %d) with slope (%d, %d) and absolute slope %f\n", a.location.X, a.location.Y, a.vector.X, a.vector.Y, a.vector.AbsoluteSlope())
		}
	}

	head := &asteroidNode{}
	var tail *asteroidNode
	cur := head
	for _, ast := range asteroidVectors {
		cur.next = &asteroidNode{ast.location, ast.vector, cur, nil}
		cur = cur.next
		tail = cur
	}

	head = head.next
	tail.next = head
	head.prev = tail

	cur = head
	for i := 1; i < n; i++ {
		cur.prev.next = cur.next
		cur.next.prev = cur.prev

		vec := cur.vector
		for cur.vector == vec {
			cur = cur.next
		}
	}

	return cur.location
}
