package util

import (
	"SIMPLPAY/model"
	"errors"
)

// utility function to get a valid user with given name
func GetAValidUserWithGivenName(userName string, userList map[string]model.User) (model.User, error){

	for key := range userList {
		if userList[key].Name == userName {
			return userList[key], nil
		}
	}
	return model.User{},errors.New("user with such name doesn't exist")
}

// utility function to get a valid merchant with given name
func GetAValidMerchantWithGivenName(merchantName string, merchantList map[string]model.Merchant) (model.Merchant, error){

	for key := range merchantList {
		if merchantList[key].Name == merchantName {
			return merchantList[key], nil
		}
	}
	return model.Merchant{},errors.New("merchant with such name doesn't exist")
}

// utility function to get a update merchant in merchant list map
func UpdateMerchantList(merchant *model.Merchant, merchantList map[string]model.Merchant) {
	delete(merchantList, merchant.Name)
	merchantList[merchant.Name] = *merchant
}

// utility function to get a update user in user list map
func UpdateUserList(userResponse *model.User, userList map[string]model.User){
	delete(userList, userResponse.Name)
	userList[userResponse.Name] = *userResponse
}