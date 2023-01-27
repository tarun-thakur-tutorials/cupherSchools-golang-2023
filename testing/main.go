package main

import "errors"

func Add(x, y int) (sum int) {
	sum = x + y
	return
}

func Subtract(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x

}

func Divide(x, y int) (int, error) {

	if x > 9999 || y > 9999 {
		return 0.0, errors.New("limit exceeded")
	}

	if y == 0 {
		return 0.0, errors.New("can not divide by zero")
	}

	return x / y, nil
}

func Factorial(x int) (int, error) {

	if x < 0 {
		return 0, errors.New("not possible for negative integers")
	}

	if x == 1 || x == 0 {
		return 1, nil
	}

	if x > 9 {
		return 0, errors.New("too large value")
	}

	fact := 1

	for i := x; i >= 1; i-- {
		fact *= i
	}

	return fact, nil
}
