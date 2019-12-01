package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(day int) []string {
	fileName := fmt.Sprintf("../input/day%d.txt", day)
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
