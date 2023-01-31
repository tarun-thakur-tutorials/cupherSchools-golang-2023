package models

type Merchant struct {
	UniqueName      string
	DiscountOffered float64
	// TransactionID   []uuid.UUID
	TotalDiscount float64
}

type name = string

var MerchantMap = make(map[name]Merchant)
