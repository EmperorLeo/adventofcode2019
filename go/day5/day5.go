package day5

import (
	"fmt"

	"github.com/EmperorLeo/adventofcode2019/util"
)

/*Silver - Part 1 */
func Silver() {
	ints := util.ReadIntcodeInstructions(5)
	computer := util.NewComputer(ints, 0)
	var err error
	stopChan := make(chan bool, 1)
	go computer.TypeRepeat(1, stopChan)
	go computer.Poll()

	for err == nil {
		err = computer.Next()
	}

	fmt.Printf("Program Terminated: Reason - %s\n", err.Error())
	stopChan <- true
	computer.Close()
	// Cannot close the computer when type/repeating, since the input channel will be closed while we are typing to it
}

/*Gold - Part 2 */
func Gold() {
	ints := util.ReadIntcodeInstructions(5)
	computer := util.NewComputer(ints, 0)
	var err error
	go computer.Type(5)
	go computer.Poll()

	for err == nil {
		err = computer.Next()
	}

	fmt.Printf("Program Terminated: Reason - %s\n", err.Error())
	computer.Close()
}
