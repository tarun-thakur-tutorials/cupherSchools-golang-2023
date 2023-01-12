package main

import "fmt"

func main() {
	defer fmt.Println("\ndefined at first, executed at last")
	code := 10
	switch code {
	case 1:
		// whatever happens at code=1
		fallthrough
	case 2:
		// whatever happens at code=2

	// .....
	case 10:
		fmt.Println("code is equal to 10 inside upper switch")
	default:
		fmt.Println("inside the default")
	}
	/*
	   if v:=10; v<10{
	   }
	*/
	switch code := 1; code {
	case 1:
		// whatever happens at code=1
		// fmt.Println("code is equal to 1 inside lower switch case")
		code = 3
		fallthrough
	case 2:
		// whatever happens at code=2
		// fmt.Println("code is equal to 2 inside lower switch case")
	case 3:
		// whatever happens at code=3
		fmt.Println("people saying yes won")

	// .....
	case 10:
		fmt.Println("code is equal to 10")
		break
	default:
		fmt.Println("inside the default")
	}
	fmt.Println("the gloabal value of code is ", code)

	// type switch

	var v interface{}
	// v = 10
	v = "some_string"
	// v= false
	// v = 0 + 4 i
	switch q := v.(type) { // v.(type) works only in switch statement
	case bool:
		fmt.Println("this is a boolean")
	case int:
		fmt.Println("this is an integer")
	default:
		fmt.Printf("value of type : %T", q)
	}

	//
}
