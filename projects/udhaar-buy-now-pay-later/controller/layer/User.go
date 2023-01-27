package layer

import (
	"fmt"
	"udhaar/constants"
	"udhaar/models"
	"udhaar/utils"
)

func AddUser(command []string) (newUserName string, err error) {

	if len(command) != 5 {
		err = constants.ErrInvalidCommand
		return
	}

	name, err := utils.IsValidName(command[2])

	if err != nil {
		return
	}

	email, err := utils.IsValidEmail(command[3])

	if err != nil {
		return
	}

	creditLimit, err := utils.IsValidCreditLimit(command[4])

	// if the user is already present or not

	_, ok := models.UserMap[email]
	// if the user is found htne give UserAlreadyExistError
	if ok {
		err = constants.ErrUserAlreadyExists
		return
	}

	// otherwise, make a new user, save it into the map and return its name

	newUser := models.User{
		Name:        name,
		UniqueEmail: email,
		CreditLimit: creditLimit,
	}

	models.UserMap[email] = newUser

	newUserName = newUser.Name
	return
}

func GetUsersAtCreditLimit(command []string) (listAsString string, err error) {

	if len(command) != 2 {
		err = constants.ErrInvalidCommand
		return
	}

	if len(models.UserMap) > 0 {
		for _, user := range models.UserMap {
			if user.Dues == user.CreditLimit {
				listAsString += fmt.Sprintf("\n%v %v", user.Name, user.UniqueEmail)
			}
		}
	} else {
		err = constants.ErrUserListEmpty

	}

	if len(listAsString) > 0 {
		listAsString = fmt.Sprintf(
			"The list of users at credit limit is as follows: \n%v",
			listAsString,
		)
	} else {
		listAsString = " There is mo user at limit right now"
	}

	return
}
