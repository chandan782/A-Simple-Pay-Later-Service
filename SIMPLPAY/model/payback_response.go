package model

import "github.com/google/uuid"

type PayBackResponse struct {
	TransactionID uuid.UUID
	UserName      string
	Dues          float64
}
