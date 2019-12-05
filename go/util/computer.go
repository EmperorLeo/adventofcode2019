package util

import (
	"errors"
	"strconv"
)

type Computer struct {
	Ip  int
	Mem []int
}

type Output struct {
	Out int
}

func (c *Computer) Next() (*Output, error) {
	/* Implement mathy version later */
	// posOneMode := c.ip - (c.ip / 1000) - ((c.ip / 10000) * 10000)
	// posTwoMode := c.ip
	// posThreeMode :=
	// fmt.Printf("My IP = %d\n", c.Ip)
	op := c.Mem[c.Ip]
	strOP := strconv.Itoa(op)
	posOneMode := len(strOP) > 2 && strOP[len(strOP)-3] == '1'
	posTwoMode := len(strOP) > 3 && strOP[len(strOP)-4] == '1'
	posThreeMode := len(strOP) > 4 && strOP[len(strOP)-5] == '1'
	// fmt.Printf("Modes: %v, %v, %v\n", posOneMode, posTwoMode, posThreeMode)

	opcode := op % 100

	var arg1, arg2, arg3 int
	if c.Ip+1 <= len(c.Mem)-1 {
		arg1 = c.Mem[c.Ip+1]
		// fmt.Printf("Arg1 Immediate: %d\n", arg1)
		if !posOneMode && opcode != 3 {
			arg1 = c.Mem[arg1]
		}
	}
	if c.Ip+2 <= len(c.Mem)-1 {
		// fmt.Printf("Arg2 Immediate: %d\n", arg2)
		arg2 = c.Mem[c.Ip+2]
		if !posTwoMode && opcode != 3 && opcode != 4 {
			arg2 = c.Mem[arg2]
		}
	}
	if c.Ip+3 <= len(c.Mem)-1 {
		// fmt.Printf("Arg3 Immediate: %d\n", arg3)
		arg3 = c.Mem[c.Ip+3]
		if !posThreeMode && opcode != 1 && opcode != 2 && opcode != 3 && opcode != 4 && opcode != 5 && opcode != 6 && opcode != 7 && opcode != 8 {
			arg3 = c.Mem[arg3]
		}
	}

	switch opcode {
	case 1:
		// fmt.Printf("Executing %d + %d to equal %d, placing in position %d\n", arg1, arg2, arg1+arg2, arg3)
		c.Mem[arg3] = arg1 + arg2
		c.Ip += 4
	case 2:
		c.Mem[arg3] = arg1 * arg2
		c.Ip += 4
	case 3:
		// fmt.Printf("Putting 1 into position %d", arg1)
		c.Mem[arg1] = c.requestInput()
		c.Ip += 2
	case 4:
		c.Ip += 2
		return &Output{arg1}, nil
	case 5:
		if arg1 != 0 {
			c.Ip = arg2
		} else {
			c.Ip += 3
		}
	case 6:
		if arg1 == 0 {
			c.Ip = arg2
		} else {
			c.Ip += 3
		}
	case 7:
		if arg1 < arg2 {
			c.Mem[arg3] = 1
		} else {
			c.Mem[arg3] = 0
		}
		c.Ip += 4
	case 8:
		if arg1 == arg2 {
			c.Mem[arg3] = 1
		} else {
			c.Mem[arg3] = 0
		}
		c.Ip += 4
	case 99:
		return nil, errors.New("Program halted")
	default:
		return nil, errors.New("massive error")
	}

	return nil, nil
}

func (c *Computer) requestInput() int {
	return 5
}
