package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EmperorLeo/adventofcode2019/util"
)

/*Silver - Part 1 */
func Silver() {
	nums := strings.Split(util.ReadLines(7)[0], ",")
	amps := make([]*util.Computer, 5)
	var maxSignal int
	combos := getPossibleCombos(0)
	for _, combo := range combos {
		// remake the amps for each combo, i dont really want to deal with resetting instructions, input, and output
		for i := 0; i < 5; i++ {
			amps[i] = makeComputer(nums)
			defer amps[i].Close()
			amps[i].Type(combo[i])
			go amps[i].Run()
		}

		var prevAmp *util.Computer
		for _, curAmp := range amps {
			go provideInput(curAmp, prevAmp, 0)
			prevAmp = curAmp
		}

		// Get the signal from the last amp and compare it to the max
		out, ok := amps[4].Read()

		if ok && out > maxSignal {
			maxSignal = out
		}
	}

	fmt.Printf("Max possible signal: %d \n", maxSignal)
}

/*Gold - Part 2 */
func Gold() {
	nums := strings.Split(util.ReadLines(7)[0], ",")
	amps := make([]*util.Computer, 5)
	var maxSignal int
	combos := getPossibleCombos(5)

	for _, combo := range combos {
		// remake the amps for each combo, i dont really want to deal with resetting instructions, input, and output
		for i := 0; i < 5; i++ {
			amps[i] = makeComputer(nums)
			amps[i].Type(combo[i])
			// cant defer closing here cause this is the main thread and i could accidentally write to a closed channel
			// I'm too lazy to make ANOTHER goroutine just to close this guy
			go amps[i].Run()
		}

		var out, outTemp int
		ok := true
		for ok {
			var prevAmp *util.Computer
			for _, curAmp := range amps {
				go provideInput(curAmp, prevAmp, out)
				prevAmp = curAmp
			}

			outTemp, ok = amps[4].Read()
			if ok {
				out = outTemp
			}
		}

		if out > maxSignal {
			maxSignal = out
		}
	}

	fmt.Printf("Max possible signal in feedback loop: %d \n", maxSignal)
}

func provideInput(amp, prevAmp *util.Computer, initial int) {
	if prevAmp != nil {
		out, ok := prevAmp.Read()
		if !ok {
			return
		}
		amp.Type(out)
	} else {
		amp.Type(initial)
	}
}

func makeComputer(input []string) *util.Computer {
	ints := make([]int, len(input))
	for n := range input {
		i, _ := strconv.Atoi(input[n])
		ints[n] = i
	}
	computer := &util.Computer{}
	computer.LoadInstructions(ints)
	return computer
}

func getPossibleCombos(offset int) [][]int {
	combos := [][]int{}
	for a := 0 + offset; a < 5+offset; a++ {
		for b := 0 + offset; b < 5+offset; b++ {
			for c := 0 + offset; c < 5+offset; c++ {
				for d := 0 + offset; d < 5+offset; d++ {
					for e := 0 + offset; e < 5+offset; e++ {
						tester := map[int]bool{}
						tester[a] = true
						tester[b] = true
						tester[c] = true
						tester[d] = true
						tester[e] = true
						if len(tester) == 5 {
							combos = append(combos, []int{a, b, c, d, e})
						}
					}
				}
			}
		}
	}
	return combos
}
