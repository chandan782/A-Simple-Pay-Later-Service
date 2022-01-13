package service

import (
	"SIMPLPAY/model"
	"SIMPLPAY/util"
	"errors"
	"github.com/google/uuid"
)

type MerchantService interface {
	MerchantOnBoard(string, string, float64, map[string]model.Merchant) (model.Merchant, error)
	UpdateMerchant(string, float64, map[string]model.Merchant) (model.Merchant, error)
	MerchantDiscountReport(string, map[string]model.Merchant, map[uuid.UUID]model.Transaction) (model.MerchantDiscount, error)
}

// merchant onboard service
func MerchantOnBoard(name string, email string, discount float64, merchantList map[string]model.Merchant) (model.Merchant, error) {
	var newMerchant model.Merchant
	_, found := merchantList[name]
	if found {
		return newMerchant, errors.New("merchant with same name already exist")
	} else {
		newMerchant = model.Merchant{
			Name:            name,
			Email:           email,
			DiscountOffered: discount,
		}
		merchantList[newMerchant.Name] = newMerchant
	}
	return newMerchant, nil
}

// update merchant discount rate service
func UpdateMerchant(name string, newDiscountRate float64, merchantList map[string]model.Merchant) (model.Merchant, error) {
	var response model.Merchant
	merchant, err := util.GetAValidMerchantWithGivenName(name, merchantList)
	if err != nil {
		return model.Merchant{}, err
	}
	merchant.DiscountOffered = newDiscountRate
	util.UpdateMerchantList(&merchant, merchantList)
	return response, nil
}

// report merchant discount value service
func MerchantDiscountReport(name string, merchantList map[string]model.Merchant, allTransactionDetails map[uuid.UUID]model.Transaction) (model.MerchantDiscount, error) {
	var response model.MerchantDiscount

	merchant, err := util.GetAValidMerchantWithGivenName(name, merchantList)
	if err != nil {
		return response, err
	}

	var totalDiscount float64
	for i := range merchant.TransactionID {
		transactions := allTransactionDetails[merchant.TransactionID[i]]
		discount := transactions.DiscountOffered
		amount := transactions.Amount
		totalDiscount = totalDiscount + (amount * discount * 0.01)
	}

	response.TotalDiscount = totalDiscount
	return response, nil
}
