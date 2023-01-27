package main

import (
	"fmt"
)

func main() {
	/*
	   	greetingMessage := HelloName("tarun")
	   	fmt.Println(greetingMessage)
	   	fmt.Println(HelloName("tarun"))

	   	v := CanFastAttack(true)
	   	fmt.Println(v)

	   	kindOfTriangle, err := KindFromSides(0, 0, 0)

	   	if err != nil {
	   		fmt.Println(err)
	   	}

	   	fmt.Println(kindOfTriangle)

	   	y := GetVarNumParamRetrunList(1, 2, 3, 4, 5, 6)
	   	fmt.Println(y)

	   	input := []int{1, 2, 3, 4}
	   	z := GetListReturnList(input)
	   	fmt.Println(z)
	   }

	   func GetVarNumParamRetrunList(x ...int) (y []int) {

	   	for _, val := range x {
	   		y = append(y, val)
	   	}
	   	return
	   }

	   func GetListReturnList(x []int) (y []int) {

	   	for _, val := range x {
	   		y = append(y, val)
	   	}
	   	return
	*/

	/*SeedWithTime()

	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(100))
	}
	*/
	unitsScale := Units()
	bill := NewBill()
	ok := AddItem(bill, unitsScale, "carrot", "dozen")
	if ok {
		fmt.Println(bill)
	}

	fmt.Println("removeing some carrots")

	ok = RemoveItem(bill, unitsScale, "carrot", "half_of_a_dozen")
	if ok {
		fmt.Println(bill)
	}

}
