package day6

import (
	"fmt"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

type body struct {
	name   string
	orbits *body
	moons  []*body
}

type pair struct {
	name string
	dist int
}

/*Silver - Part 1 */
func Silver() {
	orbitMap := makeOrbitMap()

	var numOrbits int
	for _, v := range orbitMap {
		numOrbits += v.countOrbits()
	}

	fmt.Printf("%d total orbits.\n", numOrbits)

}

/*Gold - Part 2 */
func Gold() {
	orbitMap := makeOrbitMap()
	visited := map[string]bool{}
	visited["YOU"] = true
	queue := []pair{pair{"YOU", 0}}
	// BFS
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.name == "SAN" {
			fmt.Printf("%d orbital transfers required.\n", cur.dist-2)
			break
		}
		moons := orbitMap[cur.name].moons
		orbits := orbitMap[cur.name].orbits
		for _, moon := range moons {
			queue = append(queue, pair{moon.name, cur.dist + 1})
			visited[moon.name] = true
		}
		if orbits != nil && !visited[orbits.name] {
			queue = append(queue, pair{orbits.name, cur.dist + 1})
			visited[orbits.name] = true
		}
	}
}

func makeOrbitMap() map[string]*body {
	input := util.ReadLines(6)
	orbitMap := map[string]*body{}
	for _, i := range input {
		relationship := strings.Split(i, ")")
		var planet, moon *body
		if orbitMap[relationship[0]] != nil {
			planet = orbitMap[relationship[0]]
		} else {
			planet = &body{relationship[0], nil, []*body{}}
			orbitMap[planet.name] = planet
		}
		if orbitMap[relationship[1]] != nil {
			moon = orbitMap[relationship[1]]
			moon.orbits = planet
		} else {
			moon = &body{relationship[1], planet, nil}
			orbitMap[moon.name] = moon
		}
		planet.moons = append(planet.moons, moon)
	}
	return orbitMap
}

func (b *body) countOrbits() int {
	var numOrbits int
	for b != nil {
		numOrbits++
		b = b.orbits
	}
	return numOrbits - 1
}
