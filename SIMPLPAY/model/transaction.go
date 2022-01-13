package model

import "github.com/google/uuid"

type Transaction struct {
	Id                      uuid.UUID
	Type                    string
	SourceName              string
	DestinationName         string
	Amount                  float64
	DiscountOffered         float64
	ActualTransactionAmount float64
	DiscountedAmount        float64
	Status                  string
}
