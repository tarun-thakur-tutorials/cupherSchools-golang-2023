package main

import "errors"

// a,b,c are the sides of the triangle
// a + b > c
// a >=0 ,b >=0, c>=0

const (
	Not_A_Triangle = 1
	Equilateral    = 2
	Isosceles      = 3
	Scalene        = 4
)

func KindFromSides(a, b, c int) (int, error) {
	if (a+b <= c || b+c <= a || a+c <= b) || (a <= 0 || b <= 0 || c <= 0) {
		return Not_A_Triangle, errors.New("This is not a triangle")
	} else if a == b && b == c {
		return Equilateral, nil
	} else if (a == b) || (a == c) || (b == c) {
		return Isosceles, nil
	} else {
		return Scalene, nil
	}
}
