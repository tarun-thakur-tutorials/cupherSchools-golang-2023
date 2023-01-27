package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"udhaar/constants"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(
		constants.UdhaarAppSymbol + " \nwelcome to the Udhaar - Buy Now Pay Later application \nPlease enter your input: ",
	)

	input, err := reader.ReadString('\n')

	input = strings.Replace(input, "\n", "", -1)
	input = strings.Trim(input, " ")
	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	if _, ok := constants.ExitCommand[input]; ok {
		fmt.Println(constants.Red + "Program Terminated" + constants.White)
		fmt.Println(
			constants.Reset + "Thank You for Exploring Pay-Later Terminal Client" + constants.White,
		)
		os.Exit(0)
	}

	fmt.Println("the input provided without color is ", input)
}
