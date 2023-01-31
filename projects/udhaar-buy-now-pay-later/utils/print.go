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
			msg = fmt.Sprintf(constants.Green+" User Added: %v "+constants.White, arg)
		case constants.AddMerchant:
			msg = fmt.Sprintf(constants.Green+" Merchant Added: %v "+constants.White, arg)
		case constants.UpdateDiscount:
			msg = fmt.Sprintf(constants.Green+" Merchant Dicount: %v "+constants.White, arg)
		case constants.Txn:
			msg = fmt.Sprintf(constants.Green+"Transaction Complete. Txn ID: %v"+constants.White, arg)
		case constants.Payback:
			msg = fmt.Sprintf(constants.Green+" Payment Successful. Balance: %v"+constants.White, arg)
		case constants.GetUserDues:
			msg = fmt.Sprintf(constants.Green+" the amount due on this user: %v "+constants.White, arg)
		case constants.GetUsersAtLimit:
			msg = fmt.Sprint(constants.Green + arg + constants.White)
		case constants.GetTotalDues:
			msg = fmt.Sprintf(constants.Green+"the total due on users is: %v"+constants.White, arg)
		default:
			msg = fmt.Sprint(constants.Red + "Invalid operation code" + constants.White)
		}
		fmt.Printf("%v \n\n", msg)

	}

}
