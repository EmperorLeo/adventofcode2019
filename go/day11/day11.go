package day11

import (
	"fmt"

	"github.com/EmperorLeo/adventofcode2019/util"
)

type panel struct {
	color, paintedTimes int
}

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

/*Silver - Part 1*/
func Silver() {
	input := util.ReadIntcodeInstructions(11)

	panelMap := getPaintedPanels(util.NewComputer(input, 0), 0)
	count := len(panelMap)
	fmt.Printf("Intcode robot painted %d panels \n", count)

	grid := mapPanels(panelMap)

	fmt.Println("What the hell is this shit?")
	gridPrinter(grid)
}

/*Gold - Part 2*/
func Gold() {
	input := util.ReadIntcodeInstructions(11)

	panelMap := getPaintedPanels(util.NewComputer(input, 0), 1)

	grid := mapPanels(panelMap)

	fmt.Println("Registration identifier printed below")
	gridPrinter(grid)
}

func getPaintedPanels(computer util.IComputer, startingColor int) map[util.Coord]panel {

	directionLeftMap := map[int]int{
		north: west,
		east:  north,
		south: east,
		west:  south,
	}
	directionRightMap := map[int]int{
		north: east,
		east:  south,
		south: west,
		west:  north,
	}
	directionMap := map[int]map[int]int{
		0: directionLeftMap,
		1: directionRightMap,
	}

	curCoord := util.Coord{0, 0}
	curDirection := north

	panelMap := map[util.Coord]panel{}
	panelMap[curCoord] = panel{startingColor, 0}

	go computer.Run()

	for {
		panel := panelMap[curCoord]

		computer.Type(panel.color)

		var ok bool
		var newColor, rotation int

		newColor, ok = computer.Read()
		rotation, ok = computer.Read()
		if !ok {
			// computer halted
			break
		}

		panel.color = newColor
		panel.paintedTimes++
		panelMap[curCoord] = panel

		curDirection = directionMap[rotation][curDirection]

		switch curDirection {
		case north:
			curCoord.Y--
		case east:
			curCoord.X++
		case south:
			curCoord.Y++
		case west:
			curCoord.X--
		default:
			panic("RIP intcode")
		}
	}

	return panelMap
}

func mapPanels(panels map[util.Coord]panel) [][]int {
	var minX, maxX, minY, maxY int

	for c := range panels {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	var xOffset, yOffset int
	if minX < 0 {
		xOffset = -minX
	}
	if minY < 0 {
		yOffset = -minY
	}
	maxX++
	maxY++

	grid := make([][]int, maxY+yOffset)
	for i := range grid {
		grid[i] = make([]int, maxX+xOffset)
	}

	for c, p := range panels {
		grid[c.Y+yOffset][c.X+xOffset] = p.color
	}

	return grid
}

func gridPrinter(grid [][]int) {
	for _, r := range grid {
		for _, p := range r {
			if p == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
