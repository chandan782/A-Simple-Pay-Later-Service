package controller

import (
	"SIMPLPAY/model"
	"SIMPLPAY/service"
	"errors"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type TransactionController interface {
	Transaction(string, map[string]model.User, map[string]model.Merchant, map[string]model.Transaction) (model.Transaction, error)
	UserPayBack(string, map[string]model.User, map[string]model.Transaction) (model.PayBackResponse, error)
}

// user to merchant transaction controller
func Transaction(consoleInput string, userList map[string]model.User, merchantList map[string]model.Merchant, allTransactionDetails map[uuid.UUID]model.Transaction) (model.Transaction, error) {

	transactionData := strings.Split(consoleInput, " ")
	if len(transactionData) < 5 {
		return model.Transaction{}, errors.New("invalid input, please check and re-enter")
	}
	userName := transactionData[2]
	merchantName := transactionData[3]
	//input amount is a string
	amountStr := transactionData[4]

	//converting string credit limit to a Float value
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return model.Transaction{Status: "Failure"}, err
	}

	transactionResponse, err := service.Transaction(userName, merchantName, amount, userList, merchantList, allTransactionDetails)
	if err != nil {
		return model.Transaction{Status: "Failure"}, err
	}
	return transactionResponse, nil
}

// user to PAY-LATER transaction controller
func UserPayBack(consoleInput string, userList map[string]model.User, allTransactionDetails map[uuid.UUID]model.Transaction) (model.PayBackResponse, error) {

	transactionData := strings.Split(consoleInput, " ")
	if len(transactionData) < 3 {
		return model.PayBackResponse{}, errors.New("invalid input, please check and re-enter")
	}
	userName := transactionData[1]

	//input amount is a string
	payBackAmountStr := transactionData[2]

	//converting string credit limit to a Float value
	payBackAmount, err := strconv.ParseFloat(payBackAmountStr, 64)
	if err != nil {
		return model.PayBackResponse{}, err
	}

	userResponse, err := service.UserPayBack(userName, payBackAmount, userList, allTransactionDetails)
	if err != nil {
		return userResponse, err
	}
	return userResponse, nil
}
