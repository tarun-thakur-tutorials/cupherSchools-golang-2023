package main

import (
	"errors"
	"testing"
)

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
	{1, 2, 3, 1},
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

type InputDivide struct {
	arg1          int
	arg2          int
	expectedValue int
	expectedErr   error
}

var BunchOfInputDivide = []InputDivide{
	{1, 2, 0, nil},
	{4, 2, 2, nil},
	{2, 1, 2, nil},
	{2, 0, 0, errors.New("can not divide by zero")},
	{0, 2, 0, nil},
	{5, -1, -5, nil},
	{1000000, 1, 0, errors.New("limit exceeded")},
	{1, 1000000, 0, errors.New("limit exceeded")},
}

func TestDivide(t *testing.T) {
	for _, input := range BunchOfInputDivide {
		/*if actualValue := Add(input.arg1, input.arg2); actualValue != input.expectedValueAdd {
			t.Errorf("got %v, wanted %v", actualValue, input.expectedValueAdd)
		}*/

		actualValue, err := Divide(input.arg1, input.arg2)

		if input.expectedErr != nil && err == nil {
			t.Errorf("expected  error %v, got nil", input.expectedErr)
		} else if err != nil && input.expectedErr == nil {
			t.Errorf("expected nil error, got %v", err)
		} else if input.expectedErr != nil && err.Error() != input.expectedErr.Error() {
			t.Errorf(
				"Fail for input %v, %v got error %v, wanted error %v",
				input.arg1,
				input.arg2,
				err,
				input.expectedErr,
			)

		}
		if actualValue != input.expectedValue {
			t.Errorf(
				"Fail for input %v, %v got %v, wanted %v",
				input.arg1,
				input.arg2,
				actualValue,
				input.expectedValue,
			)
		}
	}
}

type FactorialInput struct {
	number      int
	expectedVal int
	expectedErr error
}

var FactorialTestCases = []FactorialInput{
	{0, 1, nil},
	{1, 1, nil},
	{-1, 0, errors.New("not possible for negative integers")},
	{2, 2, nil},
	{5, 120, nil},
	{17, 0, errors.New("too large value")},
	{10, 0, errors.New("too large value")},
}

func TestFactorial(t *testing.T) {
	for _, input := range FactorialTestCases {
		/*if actualValue := Add(input.arg1, input.arg2); actualValue != input.expectedValueAdd {
			t.Errorf("got %v, wanted %v", actualValue, input.expectedValueAdd)
		}*/

		actualValue, err := Factorial(input.number)

		if input.expectedErr != nil && err == nil {
			t.Errorf("expected  error %v, got nil", input.expectedErr)
		} else if err != nil && input.expectedErr == nil {
			t.Errorf("expected nil error, got %v", err)
		} else if input.expectedErr != nil && err.Error() != input.expectedErr.Error() {
			t.Errorf(
				"Fail for input %v, got error %v, wanted error %v",
				input.number,
				err,
				input.expectedErr,
			)

		}
		if actualValue != input.expectedVal {
			t.Errorf(
				"Fail for input %v, got %v, wanted %v",
				input.number,
				actualValue,
				input.expectedVal,
			)
		}
	}
}
