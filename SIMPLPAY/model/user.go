package model

import (
	"github.com/google/uuid"
)

type User struct {
	Name          string
	Email         string
	CreditLimit   float64
	Dues          float64
	TransactionID []uuid.UUID
}

type UserDuesResponse struct {
	Name	[]string
	Dues 	[]float64
}

type UserWithNilCreditLimitResponse struct {
	Name	[]string
}

