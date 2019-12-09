package day9

import (
	"fmt"

	"github.com/EmperorLeo/adventofcode2019/util"
)

/*Silver - Part 1 */
func Silver() {
	input := util.ReadIntcodeInstructions(9)
	result := operateTest(input)
	fmt.Printf("Output: %v\n", result)
}

/*Gold - Part 2 */
func Gold() {
	input := util.ReadIntcodeInstructions(9)
	result := sensorBoostMode(input)
	fmt.Printf("Output: %v\n", result)
}

func operateTest(instructions []int) []int {
	computer := util.NewComputer(instructions, 0)
	defer computer.Close()
	computer.Type(1)
	go computer.Run()
	return computer.PollResult()
}

func sensorBoostMode(instructions []int) []int {
	computer := util.NewComputer(instructions, 0)
	defer computer.Close()
	computer.Type(2)
	go computer.Run()
	return computer.PollResult()
}
