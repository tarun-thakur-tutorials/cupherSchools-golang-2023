package constants

const (
	NewUser                 = "new user"
	NewMerchant             = "new merchant"
	NewTxn                  = "new txn" // new transaction
	UpdateMerchant          = "update merchant"
	PayBack                 = "payback"
	ReportDiscount          = "report discount"
	ReportDues              = "report dues"
	ReportUserAtCreditLimit = "report users-at-credit-limit"
	ReportTotalDues         = "report total-dues"
)

var ExitCommand = map[string]bool{
	string(27): true, // escape key
	"exit":     true,
	"quit":     true,
	"stop":     true,
}

// user commands

/*
Category of commands

two words: new user, new merchant, new txn, report discount, report dues, report users-at-credit-limit, report total-dues, update merchant
one word:  payback
*/

//  business logic
// verify the command, keywords, arguments etc

// data logic
// database queries
