package day9

import (
	"reflect"
	"testing"
)

func TestOperateTest(t *testing.T) {
	tests := [][]int{
		[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
		[]int{104, 1125899906842624, 99},
	}
	expectations := [][]int{
		[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		[]int{1219070632396864},
		[]int{1125899906842624},
	}
	for i, test := range tests {
		result := operateTest(test)
		if !reflect.DeepEqual(result, expectations[i]) {
			t.Logf("Expected %v, got %v\n", expectations[i], result)
			t.Fail()
		}
	}
}

func TestSensorBoostMode(t *testing.T) {
	tests := [][]int{
		[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
		[]int{104, 1125899906842624, 99},
	}
	expectations := [][]int{
		[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		[]int{1219070632396864},
		[]int{1125899906842624},
	}
	for i, test := range tests {
		result := sensorBoostMode(test)
		if !reflect.DeepEqual(result, expectations[i]) {
			t.Logf("Expected %v, got %v\n", expectations[i], result)
			t.Fail()
		}
	}
}
