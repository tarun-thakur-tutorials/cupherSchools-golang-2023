package main

import "fmt"

func main() {

	// integer ; int , uint, int64, int32 , uint32 , uint64
	// string
	// float ; float64 float32
	// bool
	// byte
	// rune
	// complex64  4 + 3i
	// error : we have a separate data type in golang for defining errors

	// zero value : default value for the data type until any value is given to it by us
	// integer = 0
	// boolean = false
	//string =""

	// var myData int64
	//var myBoolDate bool
	// var myComplexData complex64
	// fmt.Println(myComplexData)

	//var data int =8

	/*
		i := 64
		f := float64(i)
		u := uint(f)

		strVar := "100"

		// empty var ingnores the value
		intVar, _ := strconv.Atoi(strVar)
	*/
	/*
		if err !=nil{
			log.Fatal("kills the program")
		}
	*/

	//	stringNumber := strconv.Itoa(intVar)

	// const keyword

	/*
		const Data1 = "tests"
		const data2 = 34
		const data3 = false
	*/

	// for loop with initial condition and final condition
	// and step condition

	for i := 0; i < 10; i++ {
		// do the operations
	}

	// for loop with initial condition and final condition
	// and step condition

	i := 0
	for i < 10 {
		// do the operations
		i = i + 10
	}

	// infinite loop

	for {
		// this is a infinite loop
		if i > 20 {
			break
		}
	}

	if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is greater")
	}

	if v:=10;v<10{
		fmt.Println("v is less than 10")
	}else {
		fmt.Println("v is greater")
	}








}
