package model

import "github.com/google/uuid"

type Merchant struct {
	Name            string
	Email           string
	DiscountOffered float64
	TransactionID	[]uuid.UUID
}

type MerchantDiscount struct {
	TotalDiscount	float64
}
