package day5

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

func Silver() {
	input := strings.Split(util.ReadLines(5)[0], ",")
	ints := make([]int, len(input))
	for n := range input {
		i, _ := strconv.Atoi(input[n])
		ints[n] = i
	}
	computer := util.Computer{0, ints}
	var err error
	var out *util.Output
	for err == nil {
		// ip := computer.Ip
		// fmt.Printf("BEFORE: OP = %d, ARG1 = %d, ARG2 = %d, ARG3 = %d\n", computer.Mem[ip], computer.Mem[ip+1], computer.Mem[ip+2], computer.Mem[ip+3])
		out, err = computer.Next()
		// fmt.Printf("AFTER: OP = %d, ARG1 = %d, ARG2 = %d, ARG3 = %d\n", computer.Mem[ip], computer.Mem[ip+1], computer.Mem[ip+2], computer.Mem[ip+3])
		if out != nil {
			fmt.Printf("Output: %d\n", out.Out)
		}
	}

	log.Fatal(err)
}

func Gold() {

}
