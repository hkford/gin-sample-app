package test

import (
	"app/controllers"
	"reflect"
	"testing"
)

func TestCalcNextStep(t *testing.T) {
	var result int64
	result = controllers.CalcNextStep(2)
	if result != 1 {
		t.Errorf("CalcNextStep failed. expect:%d, actual:%d", 1, result)
	}

	result = controllers.CalcNextStep(3)
	if result != 10 {
		t.Errorf("CalcNextStep failed. expect:%d, actual:%d", 10, result)
	}
}

func TestCalcCollatz(t *testing.T) {
	var result []int64
	var expected_result []int64

	result = controllers.CalcCollatz(2)
	expected_result = []int64{2, 1}
	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("CalcNextStep failed. expect:%d, actual:%d", expected_result, result)
	}

	result = controllers.CalcCollatz(6)
	expected_result = []int64{6, 3, 10, 5, 16, 8, 4, 2, 1}
	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("CalcNextStep failed. expect:%d, actual:%d", expected_result, result)
	}

	result = controllers.CalcCollatz(11)
	expected_result = []int64{11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("CalcNextStep failed. expect:%d, actual:%d", expected_result, result)
	}
}

func TestGetCollatzSequence(t *testing.T) {
	var result string
	var expected_result string

	result = controllers.GetCollatzSequence(3)
	expected_result = "3 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1"
	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("CalcNextStep failed. expect:%s, actual:%s", expected_result, result)
	}

	result = controllers.GetCollatzSequence(9)
	expected_result = "9 -> 28 -> 14 -> 7 -> 22 -> 11 -> 34 -> 17 -> 52 -> 26 -> 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1"
	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("CalcNextStep failed. expect:%s, actual:%s", expected_result, result)
	}
}
