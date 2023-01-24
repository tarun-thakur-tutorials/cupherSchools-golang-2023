package main

import "testing"

// func TestXXXXX
/*
2+2
expectedValue: 4
actualValue: 5

if expectedValue != actualValue{
  give some error
}

*/

/*
func TestAdd(t *testing.T) {

	actualValue := Add(2, 2)
	expectedValue := 4

	if actualValue != expectedValue {
		t.Errorf("got %v, wanted %v", actualValue, expectedValue)
	}

}

func TestSubtract(t *testing.T) {
	actualValue := Subtract(2, 2)
	expectedValue := 4

	if actualValue != expectedValue {
		t.Errorf("got %v, wanted %v", actualValue, expectedValue)
	}

}
*/
// Table Driven Tests

type Input struct {
	arg1             int
	arg2             int
	expectedValueAdd int
	expectedValueSub int
}

var BunchOfInputs = []Input{
	{1, 2, 0, 1},
	Input{arg1: 1, arg2: 4, expectedValueAdd: 5, expectedValueSub: 3},
	Input{1, 6, 7, 6},
	{2, 6, 8, 4},
}

func TestAdd(t *testing.T) {
	for _, input := range BunchOfInputs {
		/*if actualValue := Add(input.arg1, input.arg2); actualValue != input.expectedValueAdd {
			t.Errorf("got %v, wanted %v", actualValue, input.expectedValueAdd)
		}*/

		actualValue := Add(input.arg1, input.arg2)
		if actualValue != input.expectedValueAdd {
			t.Errorf(
				"Fail for input %v, %v got %v, wanted %v",
				input.arg1,
				input.arg2,
				actualValue,
				input.expectedValueAdd,
			)
		}
	}
}
