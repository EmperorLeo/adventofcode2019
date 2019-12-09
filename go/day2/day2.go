package day2

import (
	"fmt"

	"github.com/EmperorLeo/adventofcode2019/util"
)

func Silver() {
	ans := compute(12, 2)

	fmt.Printf("Answer is %d\n", ans)
}

func Gold() {
	ans := compute(31, 46)

	fmt.Printf("Answer is %d\n", ans)
	fmt.Printf("100 * noun + verb = %d\n", 100*31+46)
}

func compute(noun, verb int) int {
	ints := util.ReadIntcodeInstructions(2)

	ints[1] = noun
	ints[2] = verb

	computer := util.NewComputer(ints, 0)
	var err error
	for err == nil {
		err = computer.Next()
	}

	computer.Close()
	return computer.GetMem(0)
}
