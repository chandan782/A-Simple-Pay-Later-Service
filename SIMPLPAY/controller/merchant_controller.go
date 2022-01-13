package controller

import (
	"SIMPLPAY/model"
	"SIMPLPAY/service"
	"errors"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

type MerchantController interface {
	MerchantOnBoard(string, map[string]model.Merchant) (model.Merchant, error)
	UpdateMerchant(string, map[string]model.Merchant) (model.Merchant, error)
	MerchantDiscountReport(string, map[string]model.Merchant, map[uuid.UUID]model.Transaction) (model.MerchantDiscount, error)
}

// merchant onboard controller
func MerchantOnBoard(consoleInput string, merchantList map[string]model.Merchant) (model.Merchant,error){

	merchantData := strings.Split(consoleInput, " ")
	if len(merchantData) < 5 {
		return model.Merchant{}, errors.New("invalid input, please check and re-enter")
	}

	name := merchantData[2]
	email := merchantData[3]
	//input discount percent is a string and trim right most % symbol
	discountPercentStr := strings.TrimRight(merchantData[4], "%")

	//converting string discount percent to a Float value
	discountPercent, err := strconv.ParseFloat(discountPercentStr, 64)
	if err != nil {
		return model.Merchant{}, err
	}

	response, err := service.MerchantOnBoard(name, email, discountPercent, merchantList)
	if err != nil{
		return response, err
	}

	return response, nil
}

// update merchant discount rate controller
func UpdateMerchant(consoleInput string, merchantList map[string]model.Merchant) (model.Merchant, error){

	merchantData := strings.Split(consoleInput, " ")
	if len(merchantData) < 4 {
		return model.Merchant{}, errors.New("invalid input, please check and re-enter")
	}
	name := merchantData[2]
	//input discount percent is a string and trim right most % symbol
	newDiscountRateStr := strings.TrimRight(merchantData[3], "%")

	//converting string discount percent to a Float value
	newDiscountRate, err := strconv.ParseFloat(newDiscountRateStr, 64)
	if err != nil {
		return model.Merchant{}, err
	}

	response, err := service.UpdateMerchant(name, newDiscountRate, merchantList)
	if err != nil{
		return response, err
	}

	return response, nil
}

// report merchant discount value controller
func MerchantDiscountReport(consoleInput string, merchantList map[string]model.Merchant, allTransactionDetails map[uuid.UUID]model.Transaction) (model.MerchantDiscount, error) {
	merchantData := strings.Split(consoleInput, " ")
	if len(merchantData) < 3 {
		return model.MerchantDiscount{}, errors.New("invalid input, please check and re-enter")
	}

	merchantName := merchantData[2]

	response, err := service.MerchantDiscountReport(merchantName, merchantList, allTransactionDetails)
	if err != nil{
		return response, err
	}

	return response, nil

}