package util

import (
	"errors"
	"fmt"
)

/*Computer represents a basic cpu simulation */
type Computer struct {
	ip            int
	mem           []int
	input, output chan int
}

/*Next executes the next instruction, and returns the next output, if any */
func (c *Computer) Next() error {
	op := c.mem[c.ip]
	posOneMode := op%1000 >= 100
	posTwoMode := op%10000 >= 1000
	posThreeMode := op%100000 >= 10000

	opcode := op % 100
	numArgs := getNumArgs(opcode)

	var arg1, arg2, arg3 int
	if numArgs >= 1 && c.ip+1 <= len(c.mem)-1 {
		arg1 = c.mem[c.ip+1]
		if !posOneMode && !alwaysImmediateMode(opcode, 1) {
			arg1 = c.mem[arg1]
		}
	}
	if numArgs >= 2 && c.ip+2 <= len(c.mem)-1 {
		arg2 = c.mem[c.ip+2]
		if !posTwoMode && !alwaysImmediateMode(opcode, 2) {
			arg2 = c.mem[arg2]
		}
	}
	if numArgs >= 3 && c.ip+3 <= len(c.mem)-1 {
		arg3 = c.mem[c.ip+3]
		if !posThreeMode && !alwaysImmediateMode(opcode, 3) {
			arg3 = c.mem[arg3]
		}
	}

	switch opcode {
	case 1:
		c.mem[arg3] = arg1 + arg2
		c.ip += 4
	case 2:
		c.mem[arg3] = arg1 * arg2
		c.ip += 4
	case 3:
		c.mem[arg1] = c.requestInput()
		c.ip += 2
	case 4:
		c.ip += 2
		c.output <- arg1
	case 5:
		if arg1 != 0 {
			c.ip = arg2
		} else {
			c.ip += 3
		}
	case 6:
		if arg1 == 0 {
			c.ip = arg2
		} else {
			c.ip += 3
		}
	case 7:
		if arg1 < arg2 {
			c.mem[arg3] = 1
		} else {
			c.mem[arg3] = 0
		}
		c.ip += 4
	case 8:
		if arg1 == arg2 {
			c.mem[arg3] = 1
		} else {
			c.mem[arg3] = 0
		}
		c.ip += 4
	case 99:
		close(c.output)
		return errors.New("program halted")
	default:
		close(c.output)
		return errors.New("massive error")
	}

	return nil
}

/*LoadInstructions - this initializes the computer with instructions */
func (c *Computer) LoadInstructions(instructions []int) {
	c.mem = instructions
	c.input = make(chan int)
	c.output = make(chan int)
}

/*Type - Please run this in another goroutine lol*/
func (c *Computer) Type(in int) {
	c.input <- in
}

/*TypeRepeat - like Type but just spams the computer with the same input*/
func (c *Computer) TypeRepeat(in int) {
	for {
		c.Type(in)
	}
}

/*Read - this would also be helpful to run in another goroutine haha */
func (c *Computer) Read() (int, bool) {
	out, ok := <-c.output
	return out, ok
}

/*Poll - reads from the output repeatedly */
func (c *Computer) Poll() {
	var out int
	ok := true
	for {
		out, ok = c.Read()
		if !ok {
			break
		}
		fmt.Printf("Output: %d\n", out)
	}
}

/*GetMem - Gets the memory value at the index*/
func (c *Computer) GetMem(at int) int {
	return c.mem[at]
}

func (c *Computer) requestInput() int {
	return <-c.input
}

/*Close the computer when you are done with it */
func (c *Computer) Close() {

	close(c.input)
}

func getNumArgs(opcode int) int {
	var numArgs int
	switch opcode {
	case 3:
		fallthrough
	case 4:
		numArgs = 1
		break
	case 5:
		fallthrough
	case 6:
		numArgs = 2
		break
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 7:
		fallthrough
	case 8:
		numArgs = 3
		break
	case 99:
	default:
	}

	return numArgs
}

func alwaysImmediateMode(opcode, argNum int) bool {
	var always bool
	switch argNum {
	case 1:
		always = opcode == 3
	case 2:
	case 3:
		always = opcode == 1 || opcode == 2 || opcode == 7 || opcode == 8
	default:
	}
	return always
}
