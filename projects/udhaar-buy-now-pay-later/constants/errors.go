package constants

import "errors"

var (
	ErrInvalidCommand      = errors.New("Invalid Command, Please Try Again...")
	ErrCreditLimitExceeded = errors.New("Credit Limit Exceeded. Aborting Transaction")
	ErrInvalidAmout        = errors.New("Invalid Amount. Amount should be greater than 0")
	ErrEmptyName           = errors.New("Name field can not be empty")
	ErrEmptyPhoneNumber    = errors.New("Phone Number field  can not be is empty")
	ErrEmptyCreditLimit    = errors.New("CreditLimit field  can not be is empty")
	ErrInvalidName         = errors.New("Name should only contain alphabets")
	ErrInvalidCreditLimit  = errors.New("Credit limit cannot be less than zero")
	ErrInvalidPhoneNumber  = errors.New("Invalid Phone Number")
	ErrInvalidDiscount     = errors.New(
		"Discount can not be less than zero and more than 100 percent",
	)
	ErrInvalidOTP            = errors.New("Your Otp doesn't match")
	ErrUserMissing           = errors.New("User not found")
	ErrMerchantMissing       = errors.New("Merchant not found")
	ErrUserAlreadyExists     = errors.New("User already exists. Plz try with a new user")
	ErrMerchantAlreadyExists = errors.New("Merchant already exists. Plz try with a new merchant")
	ErrUserListEmpty         = errors.New("User list is empty. Please register some users")
)
