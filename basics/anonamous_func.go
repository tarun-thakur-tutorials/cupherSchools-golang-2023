package main

import "fmt"

func main1() {
	func(param int) {
		fmt.Println("this is anonamous func")
	}(2)

	myVal := 2

	v := func(a, b int) int {
		return a + b + myVal
	}(1, 2)

	fmt.Println(v)
}

/*


File Input output services : RUST
Channels: Golang
third Party API: Javascript
Core: C++

gRPC: you run talk to these services like they are the machine code. Even if you want to access the code in another machine, it triggers it like it is the local machine code.

client(c#) -> server(golang)

protobuf
*/
