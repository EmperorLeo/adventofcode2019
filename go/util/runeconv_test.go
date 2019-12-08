package util

import "testing"

func TestRuneToInt(t *testing.T) {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i, digit := range digits {
		ans, err := RuneToInt(digit)
		if err != nil || i != ans {
			t.Logf("Failure: Expected %d, got %d\n", i, ans)
			t.Fail()
		}
	}
}

func TestRuneToIntFail(t *testing.T) {
	digits := []rune{'x', '%', '🙃'}
	for _, digit := range digits {
		ans, err := RuneToInt(digit)
		if err == nil {
			t.Logf("Failure: Expected error for rune %U\n, got %d instead", digit, ans)
			t.Fail()
		}
	}
}
