package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

/*Silver - Part 1 */
func Silver() {
	passRange := strings.Split(util.ReadLines(4)[0], "-")
	min, _ := strconv.Atoi(passRange[0])
	max, _ := strconv.Atoi(passRange[1])
	var count int
	for i := min; i < max+1; i++ {
		if meetsReq(strconv.Itoa(i), false) {
			count++
		}
	}

	fmt.Println(count)
}

/*Gold - Part 2 */
func Gold() {
	passRange := strings.Split(util.ReadLines(4)[0], "-")
	min, _ := strconv.Atoi(passRange[0])
	max, _ := strconv.Atoi(passRange[1])
	var count int
	for i := min; i < max+1; i++ {
		if meetsReq(strconv.Itoa(i), true) {
			count++
		}
	}

	fmt.Println(count)
}

func meetsReq(guess string, newRule bool) bool {

	var prev int

	numMap := make(map[int]int, 10)

	for _, r := range guess {
		num := int(r - '0')
		if prev > num {
			return false
		}
		numMap[num]++

		prev = num
	}

	var twos, mores int
	for _, v := range numMap {
		if v == 2 {
			twos++
		} else if v > 2 {
			mores++
		}
	}

	if newRule {
		return twos > 0
	} else {
		return twos > 0 || mores > 0
	}
}
