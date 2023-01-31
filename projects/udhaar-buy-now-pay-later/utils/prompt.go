package utils

import (
	"fmt"
	"udhaar/constants"
)

func PromptMsg() {
	fmt.Print(constants.Blue + "\nudhaar > " + constants.White)
}
