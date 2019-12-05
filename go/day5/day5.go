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
	var out *util.Output
	go computer.Type(1)
	for err == nil {
		out, err = computer.Next()
		if out != nil {
			go computer.Type(1)
			fmt.Printf("Output: %d\n", out.Out)
		}
	}

	fmt.Printf("Program Terminated: Reason - %s\n", err.Error())
}

/*Gold - Part 1 */
func Gold() {
	computer := makeComputer()
	var err error
	var out *util.Output
	go computer.Type(5)
	for err == nil {
		out, err = computer.Next()
		if out != nil {
			fmt.Printf("Output: %d\n", out.Out)
		}
	}

	fmt.Printf("Program Terminated: Reason - %s\n", err.Error())
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
