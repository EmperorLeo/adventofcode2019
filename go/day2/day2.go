package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

func Silver() string {
	ans := computer(12, 2)

	return fmt.Sprintf("Answer is %d", ans)
}

func Gold() string {
	ans := computer(31, 46)

	fmt.Printf("Answer is %d\n", ans)
	return fmt.Sprintf("100 * noun + verb = %d", 100*31+46)
}

func computer(noun, verb int) int {
	input := util.ReadLines(2)[0]
	strInts := strings.Split(input, ",")
	ints := make([]int, len(strInts))
	for n := range strInts {
		i, _ := strconv.Atoi(strInts[n])
		ints[n] = i
	}

	ints[1] = noun
	ints[2] = verb

	for n := 0; n < len(ints); {
		if ints[n] == 99 {
			break
		} else {
			opcode, op1, op2, loc := ints[n], ints[ints[n+1]], ints[ints[n+2]], ints[n+3]

			if opcode == 1 {
				ints[loc] = op1 + op2
			} else if opcode == 2 {
				ints[loc] = op1 * op2
			}
		}

		n += 4
	}

	return ints[0]
}
