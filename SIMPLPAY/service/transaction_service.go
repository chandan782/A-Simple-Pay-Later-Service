package service

import (
	"SIMPLPAY/constants"
	"SIMPLPAY/model"
	"SIMPLPAY/util"
	"errors"
	"github.com/google/uuid"
)

type TransactionService interface {
	Transaction(string, string, float64, map[string]model.User, map[string]model.Merchant, map[uuid.UUID]model.Merchant) (model.Transaction, error)
	UserPayBack(string, float64, map[string]model.User, map[uuid.UUID]model.Transaction) (model.PayBackResponse, error)
}

// user to merchant transaction service
func Transaction(userName string, merchantName string, amount float64, userList map[string]model.User, merchantList map[string]model.Merchant, allTransactionDetails map[uuid.UUID]model.Transaction) (model.Transaction, error) {

	var transactionResponse model.Transaction
	userResponse, err := util.GetAValidUserWithGivenName(userName, userList)
	if err != nil {
		return model.Transaction{Status: "Failure"}, err
	}
	merchantResponse, err := util.GetAValidMerchantWithGivenName(merchantName, merchantList)
	if err != nil {
		return model.Transaction{Status: "Failure"}, err
	}

	discountedAmount := merchantResponse.DiscountOffered * amount * 0.01
	amountToBeTransferred := amount - discountedAmount

	if userResponse.CreditLimit >= amountToBeTransferred {
		transactionId, err := uuid.NewUUID()
		if err != nil {
			return model.Transaction{Status: "Failure"}, err
		}

		transactionResponse = model.Transaction{
			Id:                      transactionId,
			Type:                    constants.UserToMerchantTransaction,
			SourceName:              userName,
			DestinationName:         merchantName,
			Amount:                  amount,
			DiscountOffered:         merchantResponse.DiscountOffered,
			ActualTransactionAmount: amountToBeTransferred,
			DiscountedAmount:        discountedAmount,
			Status:                  "Success",
		}
		userResponse.CreditLimit = userResponse.CreditLimit - transactionResponse.Amount
		userResponse.Dues = userResponse.Dues + transactionResponse.Amount

		userResponse.TransactionID = append(userResponse.TransactionID, transactionId)
		merchantResponse.TransactionID = append(merchantResponse.TransactionID, transactionId)

		allTransactionDetails[transactionResponse.Id] = transactionResponse
		util.UpdateUserList(&userResponse, userList)
		util.UpdateMerchantList(&merchantResponse, merchantList)
	} else {
		return model.Transaction{Status: "Failure"}, errors.New("rejected! (reason: credit limit)")
	}

	return transactionResponse, nil
}

// user to SIMPLPAY transaction service
func UserPayBack(userName string, paybackAmount float64, userList map[string]model.User, allTransactionDetails map[uuid.UUID]model.Transaction) (model.PayBackResponse, error) {

	var paybackResponse model.PayBackResponse

	userResponse, err := util.GetAValidUserWithGivenName(userName, userList)
	if err != nil {
		return model.PayBackResponse{}, err
	}

	if userResponse.Dues < paybackAmount {
		return model.PayBackResponse{}, errors.New("payback amount is more than user dues")
	}

	if userResponse.Dues >= paybackAmount {
		transactionId, err := uuid.NewUUID()
		if err != nil {
			return model.PayBackResponse{}, err
		}

		userResponse.Dues = userResponse.Dues - paybackAmount
		userResponse.CreditLimit = userResponse.CreditLimit + paybackAmount

		paybackResponse.TransactionID = transactionId
		paybackResponse.Dues = userResponse.Dues
		paybackResponse.UserName = userResponse.Name

		transactionResponse := model.Transaction{
			Id:              transactionId,
			Type:            constants.UserToPayBackTransaction,
			SourceName:      userName,
			DestinationName: constants.PayBackDestinationNameForUser,
			Amount:          paybackAmount,
			DiscountOffered: 0.0,
			Status:          "Success",
		}
		util.UpdateUserList(&userResponse, userList)
		allTransactionDetails[transactionId] = transactionResponse
	}
	return paybackResponse, nil
}
