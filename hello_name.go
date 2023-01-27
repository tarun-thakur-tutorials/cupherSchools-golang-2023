package main

import "fmt"

/*
function : HelloName ()
input: Name string
output: "Hello "+ Name
*/

func HelloName(name string) (greetingMessage string) {
	greetingMessage = fmt.Sprintf("hello %v", name) // "hello " +name
	return
}
