package day9

import (
	"github.com/EmperorLeo/adventofcode2019/util"
)

/*Silver - Part 1 */
func Silver() {
	input := util.ReadIntcodeInstructions(9)
	operateTest(input)
}

/*Gold - Part 2 */
func Gold() {
	input := util.ReadIntcodeInstructions(9)
	sensorBoostMode(input)
}

func operateTest(instructions []int) {
	computer := util.NewComputer(instructions, 0)
	computer.Type(1)
	go computer.Run()
	computer.Poll()
	computer.Close()
}

func sensorBoostMode(instructions []int) {
	computer := util.NewComputer(instructions, 0)
	computer.Type(2)
	go computer.Run()
	computer.Poll()
	computer.Close()
}
