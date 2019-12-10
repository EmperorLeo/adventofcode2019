package day10

import (
	"testing"

	"github.com/EmperorLeo/adventofcode2019/util"
)

func TestGetBestStationPosition(t *testing.T) {
	tests := [][]string{
		[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		},
		// []string{
		// 	"#.........",
		// 	"...#......",
		// 	"...#..#...",
		// 	".####....#",
		// 	"..#.#.#...",
		// 	".....#....",
		// 	"..###.#.##",
		// 	".......#..",
		// 	"....#...#.",
		// 	"...#..#..#",
		// },
		[]string{
			"......#.#.",
			"#..#.#....",
			"..#######.",
			".#.#.###..",
			".#..#.....",
			"..#....#.#",
			"#..#....#.",
			".##.#..###",
			"##...#..#.",
			".#....####",
		},
		[]string{
			"#.#...#.#.",
			".###....#.",
			".#....#...",
			"##.#.#.#.#",
			"....#.#.#.",
			".##..###.#",
			"..#...##..",
			"..##....##",
			"......#...",
			".####.###.",
		},
		[]string{
			".#..#..###",
			"####.###.#",
			"....###.#.",
			"..###.##.#",
			"##.##.#.#.",
			"....###..#",
			"..#.#..#.#",
			"#..#.#.###",
			".##...##.#",
			".....#.#..",
		},
		[]string{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		},
	}
	expectations := []util.Coord{
		util.Coord{3, 4},
		// util.Coord{}
		util.Coord{5, 8},
		util.Coord{1, 2},
		util.Coord{6, 3},
		util.Coord{11, 13},
	}
	expectedCounts := []int{8, 33, 35, 41, 210}

	for i, test := range tests {
		ans, count, _ := getBestStationPosition(test)
		exp := expectations[i]
		expCount := expectedCounts[i]
		if ans != exp || expCount != count {
			t.Logf("Expected (%d, %d) and count %d, got (%d, %d) and count %d.\n", exp.X, exp.Y, expCount, ans.X, ans.Y, count)
			t.Fail()
		}
	}
}

func TestGetNthAsteroid(t *testing.T) {
	tests := [][]string{
		[]string{
			".#....#####...#..",
			"##...##.#####..##",
			"##...#...#.#####.",
			"..#.....#...###..",
			"..#.#.....#....##",
		},
		[]string{
			"#####",
			"#####",
			"#####",
			"#####",
			"#####",
		},
		[]string{
			"#####",
			"#####",
			"#####",
			"#####",
			"#####",
		},
		[]string{
			"#####",
			"#.###",
			"#####",
			"#####",
			"#####",
		},
	}
	numAsteroids := []int{9, 1, 24, 15}
	stations := []util.Coord{
		util.Coord{8, 3},
		util.Coord{2, 2},
		util.Coord{2, 2},
		util.Coord{2, 2},
	}
	expectations := []util.Coord{
		util.Coord{15, 1},
		util.Coord{2, 1},
		util.Coord{0, 0},
		util.Coord{0, 0},
	}

	for i, test := range tests {
		exp := expectations[i]
		_, _, asteroids := getBestStationPosition(test)
		result := getNthAsteroid(stations[i], asteroids, numAsteroids[i])
		if exp != result {
			t.Logf("Expected (%d, %d), got (%d, %d).\n", exp.X, exp.Y, result.X, result.Y)
			t.Fail()
		}
	}
}
