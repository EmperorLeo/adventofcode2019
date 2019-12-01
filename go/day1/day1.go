package day1

import (
	"fmt"
	"strconv"

	"github.com/EmperorLeo/adventofcode2019/util"
)

func Silver() string {
	return fmt.Sprintf("Requirement = %d", getFuelReq())
}

func Gold() string {
	return fmt.Sprintf("Requirement = %d", getRealFuelReq())
}

func getFuelReq() int {
	modules := util.ReadLines(1)
	var requirement int
	for _, module := range modules {
		val, _ := strconv.Atoi(module)
		requirement += ((val / 3) - 2)
	}
	return requirement
}

func getRealFuelReq() int {
	modules := util.ReadLines(1)
	var requirement int
	for _, module := range modules {
		val, _ := strconv.Atoi(module)
		for val > 0 {
			fuelNeeded := ((val / 3) - 2)
			if fuelNeeded < 0 {
				fuelNeeded = 0
			}
			requirement += fuelNeeded
			val = fuelNeeded
		}
	}

	return requirement
}
