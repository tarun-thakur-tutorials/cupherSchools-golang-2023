package main

import (
	"fmt"
) 

func main() {
	/*
	   //fmt.Print("Hello world\n")
	   fmt.Println("Hello world")
	   // Print("hello ")

	   name := "tarun"
	   fmt.Printf("hello %v", name)
	   // %d -> int
	   // %s -> string
	   // %w -> error
	   // %v -> any data
	   // %f -> float
	   name1 := "tarun"
	   name2 := "rohan"
	   name3 := "kartick"
	   name4 := "somraj"
	   fmt.Printf("\nhello %v %v %v %v\n", name1, name2, name3, name4)
	   fmt.Printf("\nhello %[1]v %[1]v %[1]v %[1]v\n", name1)

	   var name5 string = "tarun"
	   fmt.Print(name5)

	   name6 := "tarun"

	   fmt.Println(name6)
	*/
	greetingMessage := Greeting("tarun")
	fmt.Println(greetingMessage)

	v1 := GreetingWithNamedOutputParam("rahul")

	fmt.Println(v1)
	v1="asdasd"

	

	

}

func Print(input string) {
	fmt.Println(input)
}

func Greeting(personName string) string {
	return "hello " + personName
}

func GreetingWithNamedOutputParam(personName string) (greetingMessage string) {

	if personName == "tarun" {
		greetingMessage = " hello sir"
		return
	}

	greetingMessage = "hello " + personName
	return
}

//GreetingWithVarNumOfParam("tarun","rahul","kanit","asad","asdasd")

func GreetingWithVarNumOfParam(personNames ...string) (greetingMessages []string){
	

	return
}
