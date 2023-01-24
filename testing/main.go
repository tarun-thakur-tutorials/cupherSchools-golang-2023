package main

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
