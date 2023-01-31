package models

type User struct {
	Name              string
	UniquePhoneNumber string
	CreditLimit       float64
	Dues              float64
	// TransactionID     []uuid.UUID
}

type phoneNumber = string

var UserMap = make(map[phoneNumber]User)
