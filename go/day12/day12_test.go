package day12

import "testing"

func TestSystem(t *testing.T) {

	tests := [][]string{
		[]string{
			"<x=-1, y=0, z=2>",
			"<x=2, y=-10, z=-7>",
			"<x=4, y=-8, z=8>",
			"<x=3, y=5, z=-1>",
		},
		[]string{
			"<x=-8, y=-10, z=0>",
			"<x=5, y=5, z=10>",
			"<x=2, y=-7, z=3>",
			"<x=9, y=-8, z=-3>",
		},
	}
	runs := []int{10, 100}
	expectations := []int{179, 1940}

	for i, test := range tests {
		expected := expectations[i]
		moons := parseInput(test)
		advanceSteps(moons, runs[i])
		totalEnergy := calculateEnergy(moons)
		if expected != totalEnergy {
			t.Logf("Expected %d, got %d\n", expected, totalEnergy)
			t.Fail()
		}
	}
}
