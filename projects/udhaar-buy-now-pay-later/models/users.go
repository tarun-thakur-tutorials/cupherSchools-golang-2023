package models

import "github.com/google/uuid"

type User struct {
	Name          string
	UniqueEmail   string
	CreditLimit   float64
	Dues          float64
	TransactionID []uuid.UUID
}

type email = string

var UserMap = make(map[email]User)
