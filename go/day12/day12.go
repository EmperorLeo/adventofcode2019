package day12

import (
	"fmt"
	"math"

	"github.com/EmperorLeo/adventofcode2019/util"
)

const debug bool = false

type eightPair struct {
	a, b, c, d, e, f, g, h int
}

/*Silver - Part 1*/
func Silver() {
	input := util.ReadLines(12)

	moons := parseInput(input)

	advanceSteps(moons, 1000)

	totalEnergy := calculateEnergy(moons)

	fmt.Printf("The total energy in the system is %d\n", totalEnergy)
}

/*Gold - Part 2*/
func Gold() {
	input := util.ReadLines(12)

	moons := parseInput(input)

	periodX, periodY, periodZ, offsetX, offsetY, offsetZ := advanceSteps(moons, 1000000)
	fmt.Printf("PX = %d, PY = %d, PZ = %d, OX = %d, OY = %d, OZ = %d\n", periodX, periodY, periodZ, offsetX, offsetY, offsetZ)
}

func parseInput(lines []string) []*util.PosAndVector3 {

	moons := make([]*util.PosAndVector3, 4)

	if debug {
		fmt.Printf("After 0 steps:\n")
	}

	for i, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)

		if debug {
			fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", x, y, z, 0, 0, 0)
		}

		moons[i] = &util.PosAndVector3{
			&util.Coord3{x, y, z},
			&util.Coord3{},
		}
	}

	if debug {
		fmt.Println()
	}

	return moons
}

func advanceSteps(moons []*util.PosAndVector3, steps int) (int, int, int, int, int, int) {

	xMap := map[eightPair]int{}
	startX := math.MaxInt64
	yMap := map[eightPair]int{}
	startY := math.MaxInt64
	zMap := map[eightPair]int{}
	startZ := math.MaxInt64
	var xPeriod, yPeriod, zPeriod int

	for i := 0; i < steps; i++ {

		pairX := eightPair{
			moons[0].Pos.X,
			moons[1].Pos.X,
			moons[2].Pos.X,
			moons[3].Pos.X,
			moons[0].Vec.X,
			moons[1].Vec.X,
			moons[2].Vec.X,
			moons[3].Vec.X,
		}
		existing, exists := xMap[pairX]
		if exists {
			xPeriod = i - existing
			if i < startX {
				startX = i
			}
		}
		xMap[pairX] = i

		pairY := eightPair{
			moons[0].Pos.Y,
			moons[1].Pos.Y,
			moons[2].Pos.Y,
			moons[3].Pos.Y,
			moons[0].Vec.Y,
			moons[1].Vec.Y,
			moons[2].Vec.Y,
			moons[3].Vec.Y,
		}
		existing, exists = yMap[pairY]
		if exists {
			yPeriod = i - existing
			if i < startY {
				startY = i
			}
		}
		yMap[pairY] = i

		pairZ := eightPair{
			moons[0].Pos.Z,
			moons[1].Pos.Z,
			moons[2].Pos.Z,
			moons[3].Pos.Z,
			moons[0].Vec.Z,
			moons[1].Vec.Z,
			moons[2].Vec.Z,
			moons[3].Vec.Z,
		}
		existing, exists = zMap[pairZ]
		if exists {
			zPeriod = i - existing
			if i < startZ {
				startZ = i
			}
		}
		zMap[pairZ] = i

		if xPeriod > 0 && yPeriod > 0 && zPeriod > 0 {
			break
		}

		for j := 0; j < len(moons)-1; j++ {
			for k := j + 1; k < len(moons); k++ {
				m1 := moons[j]
				m2 := moons[k]

				dx1, dx2 := pullTogether(m1.Pos.X, m2.Pos.X)
				m1.Vec.X += dx1
				m2.Vec.X += dx2

				dy1, dy2 := pullTogether(m1.Pos.Y, m2.Pos.Y)
				m1.Vec.Y += dy1
				m2.Vec.Y += dy2

				dz1, dz2 := pullTogether(m1.Pos.Z, m2.Pos.Z)
				m1.Vec.Z += dz1
				m2.Vec.Z += dz2
			}
		}

		if debug {
			fmt.Printf("After %d steps:\n", i+1)
		}

		for _, moon := range moons {
			moon.Pos.X += moon.Vec.X
			moon.Pos.Y += moon.Vec.Y
			moon.Pos.Z += moon.Vec.Z

			if debug {
				fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", moon.Pos.X, moon.Pos.Y, moon.Pos.Z, moon.Vec.X, moon.Vec.Y, moon.Vec.Z)
			}
		}

		if debug {
			fmt.Println()
		}
	}

	return xPeriod, yPeriod, zPeriod, startX, startY, startZ
}

func calculateEnergy(moons []*util.PosAndVector3) int {
	var total int
	for _, moon := range moons {
		potential := util.Abs(moon.Pos.X) + util.Abs(moon.Pos.Y) + util.Abs(moon.Pos.Z)
		kinetic := util.Abs(moon.Vec.X) + util.Abs(moon.Vec.Y) + util.Abs(moon.Vec.Z)
		total += (potential * kinetic)
	}
	return total
}

func pullTogether(i, j int) (int, int) {
	if i > j {
		return -1, 1
	} else if i < j {
		return 1, -1
	}
	return 0, 0
}

func printMoons(moons []*util.PosAndVector3) {
	for _, moon := range moons {
		fmt.Printf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n", moon.Pos.X, moon.Pos.Y, moon.Pos.Z, moon.Vec.X, moon.Vec.Y, moon.Vec.Z)
	}
}
