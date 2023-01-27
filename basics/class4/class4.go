package main

import "fmt"

// golang also supports array
// but it is a popular practice to use slice instead of array
// slice is similar to array but without a fixed size
// array : var myArray [5]int
////// how to accomodate more elements to array after it is fully filled
////// you define a new array of double the size of the original
////// first, copy the content of the original array in the new array
////// now, you save the new values in the new array and delete the old array
// slice :  var mySlice []int

func main() {

	// index of array and slice starts from 0
	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArray)
	/*
		v := "some_string"

		var j string
		j = "some value"
	*/
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice)

	mySlice2 := make([]int, 10) // []int{0,0,0,0,0,0,0,0,0,0}

	fmt.Println("printing slice2 ", mySlice2)
	mySlice3 := make([]int, 0, 10000) // []int{}

	fmt.Println("printing mySlice3", mySlice3)
	mySlice4 := make([]int, 0)

	for i := 0; i < 10; i++ {
		mySlice4 = append(mySlice4, 3) // 2,4,8,16
	}

	fmt.Println("printing mySlice4", mySlice4)
	// new function returns the pointer to the object
	// but make function returns the object
	// mySlice5:=new([]int)

	slice5 := mySlice4[:10] // index 0,1,2,3,4,5,6,7,8,9

	slice5[4] = 123

	fmt.Println("slice4 after I alter an element of slice5", mySlice4)

	// in case we want to delete the index 4

	slice5 = append(slice5[:4], slice5[5:]...)
	// slice5[5:]... = slice5[5], slice5[6], slice5[7], slice5[8], slice5[9]...
	// length of slice
	l := len(slice5)
	fmt.Println(l)

	for i := 0; i < len(slice5); i++ {
		slice5[i] = slice5[i] - 5
		fmt.Println(slice5[i])
	}

	for index := range slice5 {

		slice5[index] = slice5[index] - 5
		fmt.Println(slice5[index])
	}

	for index, value := range slice5 {
		if index == 0 {
			value = 6
		}
		fmt.Println("index: ", index, " value: ", value)
	}

}
