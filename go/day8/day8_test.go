package day8

import "testing"

func TestGetFewestZeroesValue(t *testing.T) {
	t.Log("Running getFewestZeroesValueTest")
	ans := getFewestZeroesValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 3, 2)
	if ans != 1 {
		t.Logf("Expected 1, got %d\n", ans)
		t.FailNow()
	}
}

func TestGetFewestZeroesValueTwo(t *testing.T) {
	t.Log("Running getFewestZeroesValue test")
	ans := getFewestZeroesValue([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2,
	}, 5, 4)
	if ans != 30 {
		t.Logf("Expected 30, got %d\n", ans)
		t.FailNow()
	}
}

func TestDecodeImage(t *testing.T) {
	t.Log("Running decodeImage test")
	rawImageData := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}
	image := decodeImage(rawImageData, 2, 2)
	if len(image) != 2 || len(image[0]) != 2 || image[0][0] != 0 || image[0][1] != 1 || image[1][0] != 1 || image[1][1] != 0 {
		t.Logf("RIP:\n%d%d\n%d%d", image[0][0], image[0][1], image[1][0], image[1][1])
		t.FailNow()
	}
}
