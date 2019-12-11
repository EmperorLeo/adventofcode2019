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

func Silver() {
	input := util.ReadIntcodeInstructions(11)

	panelMap := getPaintedPanels(util.NewComputer(input, 0))
	count := len(panelMap)
	fmt.Printf("Intcode robot painted %d panels \n", count)
}

func Gold() {

}

func getPaintedPanels(computer *util.Computer) map[util.Coord]panel {

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
			curCoord.Y++
		case east:
			curCoord.X++
		case south:
			curCoord.Y--
		case west:
			curCoord.X--
		default:
			panic("RIP intcode")
		}
	}

	return panelMap
}
