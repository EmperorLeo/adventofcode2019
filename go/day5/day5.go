package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

/*Silver - Part 1 */
func Silver() {
	computer := makeComputer()
	var err error
	go computer.TypeRepeat(1)
	go computer.Poll()

	for err == nil {
		err = computer.Next()
	}

	fmt.Printf("Program Terminated: Reason - %s\n", err.Error())
	// Cannot close the computer when type/repeating, since the input channel will be closed while we are typing to it
}

/*Gold - Part 2 */
func Gold() {
	computer := makeComputer()
	var err error
	go computer.Type(5)
	go computer.Poll()

	for err == nil {
		err = computer.Next()
	}

	fmt.Printf("Program Terminated: Reason - %s\n", err.Error())
	computer.Close()
}

func makeComputer() *util.Computer {
	input := strings.Split(util.ReadLines(5)[0], ",")
	ints := make([]int, len(input))
	for n := range input {
		i, _ := strconv.Atoi(input[n])
		ints[n] = i
	}
	computer := &util.Computer{}
	computer.LoadInstructions(ints)
	return computer
}
