package day8

import (
	"fmt"
	"math"

	"github.com/EmperorLeo/adventofcode2019/util"
)

const (
	width  int = 25
	height int = 6
)

/*Silver - Part 1 */
func Silver() {
	input := readInput()
	ans := getFewestZeroesValue(input, width, height)
	fmt.Printf("Best layer Ones * Twos: %d\n", ans)
}

/*Gold - Part 2 */
func Gold() {
	input := readInput()
	ans := decodeImage(input, width, height)
	fmt.Println("Image below:")
	for _, r := range ans {
		for _, pix := range r {
			if pix == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getFewestZeroesValue(pixels []int, w, h int) int {
	layerSize := w * h
	totalSize := len(pixels)
	layerCount := totalSize / layerSize
	var fewestZeroCount, bestLayer int
	fewestZeroCount = math.MaxInt32
	for i := 0; i < layerCount; i++ {
		zeroCount := findCount(pixels[i*layerSize:(i+1)*layerSize], 0)
		if zeroCount < fewestZeroCount {
			fewestZeroCount = zeroCount
			bestLayer = i
		}
	}
	onesCount := findCount(pixels[bestLayer*layerSize:(bestLayer+1)*layerSize], 1)
	twosCount := findCount(pixels[bestLayer*layerSize:(bestLayer+1)*layerSize], 2)

	return onesCount * twosCount
}

func decodeImage(pixels []int, w, h int) [][]int {
	layerSize := w * h
	totalSize := len(pixels)
	layerCount := totalSize / layerSize
	decoded := make([][]int, h)
	for x := 0; x < h; x++ {
		decoded[x] = make([]int, w)
	}

	for i := layerCount - 1; i >= 0; i-- {
		layer := pixels[i*layerSize : (i+1)*layerSize]
		for x := 0; x < h; x++ {
			for y := 0; y < w; y++ {
				newPixel := layer[y+(x*w)]
				// if the pixel is transparent, don't overwrite
				if newPixel != 2 {
					decoded[x][y] = newPixel
				}
			}
		}
	}

	return decoded
}

func findCount(pixels []int, pixelNum int) int {
	var count int
	for _, p := range pixels {
		if p == pixelNum {
			count++
		}
	}
	return count
}

func readInput() []int {
	input := util.ReadLines(8)[0]
	pixels := make([]int, len(input))
	for i, pix := range input {
		pixels[i], _ = util.RuneToInt(pix)
	}
	return pixels
}
