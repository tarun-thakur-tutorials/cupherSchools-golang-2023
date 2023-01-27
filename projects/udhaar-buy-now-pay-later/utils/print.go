package utils

import (
	"fmt"
	"udhaar/constants"
)

func PrintMsg(err error, operationCode int, arg string) {
	var msg string
	if err != nil {
		fmt.Println(constants.Red + err.Error() + constants.White + "\n\n")

	} else {
		switch operationCode {
		case constants.AddUser:
			msg = fmt.Sprintf(" User Added: %v ", arg)
		default:
			msg = fmt.Sprint("Invalid operation code")
		}
		fmt.Printf("%v \n\n", msg)

	}

}
