package repo

import (
	"udhaar/constants"
	"udhaar/database"
	"udhaar/models"
)

func IsUserExists(phoneNumber string) error {
	rows, err := database.Db.Query("select * from user where unique_phone_number = ? ", phoneNumber)
	if err != nil {
		return err
	}

	if rows.Next() {
		err = constants.ErrUserAlreadyExists
		return err
	}
	return nil
}

func InsertUser(newUser models.User) (err error) {
	result, err := database.Db.Exec(
		"INSERT INTO user (name, unique_phone_number,credit_limit) VALUES (?,?,?) ",
		newUser.Name,
		newUser.UniquePhoneNumber,
		newUser.CreditLimit,
	)
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// paswword == UniquePhoneNumber
	result, err = database.Db.Exec(
		"INSERT INTO userCreds (id, pass) VALUES (?,?) ",
		id,
		newUser.UniquePhoneNumber,
	)
	return
}
