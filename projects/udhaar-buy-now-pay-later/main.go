package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"udhaar/constants"
	"udhaar/controller/layer"
	"udhaar/utils"
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

	command := strings.Split(input, " ")
	var rootCommand string

	if command[0] != "payback" {
		rootCommand = strings.Join(command[:2], " ") // index 0 , 1, separator: " "
	} else {
		rootCommand = command[0]
	}
	var arg string
	var operationCode int
	switch rootCommand {
	case constants.NewUser:
		arg, err = layer.AddUser(command)
		operationCode = constants.AddUser
	case constants.NewMerchant:
	case constants.NewTxn:
	case constants.UpdateMerchant:
	case constants.PayBack:
	case constants.ReportDiscount:
	case constants.ReportDues:
	case constants.ReportUserAtCreditLimit:
		arg, err = layer.GetUsersAtCreditLimit(command)
	case constants.ReportTotalDues:
	default:
		err = constants.ErrInvalidCommand
	}

	utils.PrintMsg(err, operationCode, arg)
}
