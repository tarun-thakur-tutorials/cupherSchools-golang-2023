package layer

import (
	"udhaar/constants"
	"udhaar/controller/repo"
	"udhaar/models"
	"udhaar/utils"

	"github.com/spf13/cast"
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

	phoneNumber, err := utils.IsValidPhoneNumber(command[3])
	if err != nil {
		return
	}
	err = repo.IsUserExists(phoneNumber)

	if err != nil {
		return
	}

	err = utils.ValidateViaWhatsapp(phoneNumber)
	if err != nil {
		err = constants.ErrInvalidOTP
		return
	}
	creditLimit, err := utils.IsValidCreditLimit(command[4])

	newUser := models.User{
		Name:              name,
		UniquePhoneNumber: phoneNumber,
		CreditLimit:       creditLimit,
	}

	// models.UserMap[phoneNumber] = newUser

	err = repo.InsertUser(newUser)

	if err != nil {
		return
	}

	newUserName = newUser.Name
	return
}

func GetUsersAtCreditLimit(command []string) (listAsString string, err error) {
	if len(command) != 2 {
		err = constants.ErrInvalidCommand
		return
	}
	/*
		if len(models.UserMap) > 0 {
			for _, user := range models.UserMap {
				if user.Dues == user.CreditLimit {
					listAsString += fmt.Sprintf("\n%v %v", user.Name, user.UniquePhoneNumber)
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
	*/
	return
}

func ShowUserDues(command []string) (duesString string, err error) {
	/*
		report dues <user-email>
	*/

	if len(command) != 3 {
		err = constants.ErrInvalidCommand
		return
	}
	/*
		userEmail := command[2]

		// check if the user/merchant exists or not

		userObject, isUserExist := models.UserMap[userEmail]

		// if the user doesn't exist
		if !isUserExist {
			err = constants.ErrUserMissing
			return
		}
		duesString = cast.ToString(userObject.Dues)
	*/
	return
}

func GetTotalDues(command []string) (totalDuesString string, err error) {
	totalDues := 0.0
	// sum	the total dues of the users

	for _, userObj := range models.UserMap {
		totalDues += userObj.Dues
	}

	totalDuesString = cast.ToString(totalDues)
	return
}
