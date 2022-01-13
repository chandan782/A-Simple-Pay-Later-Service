package service

import (
	"SIMPLPAY/model"
	"SIMPLPAY/util"
	"errors"
	"fmt"
)

type UserService interface {
	UserOnBoard(string, string, float64, map[string]model.User) (model.User, error)
	AllUserWithDues(map[string]model.User) (model.UserDuesResponse, error)
	UserDues(string, map[string]model.User) (float64, error)
	UsersExhaustedCreditLimit(map[string]model.User) (model.UserWithNilCreditLimitResponse, error)
}

// user onboard service
func UserOnBoard(name string, email string, creditLimit float64, userList map[string]model.User) (model.User, error) {

	var newUser model.User
	_, found := userList[name]
	if found {
		return newUser, errors.New("user already exist with same email id")
	} else {
		newUser = model.User{
			Name:        name,
			Email:       email,
			CreditLimit: creditLimit,
			Dues:        0.0,
		}
		userList[newUser.Name] = newUser
	}
	return newUser, nil
}

// list of user whose dues are pending service
func AllUserWithDues(userList map[string]model.User) (model.UserDuesResponse, error) {

	var response model.UserDuesResponse
	for _, values := range userList {
		if int(values.Dues) != 0 {
			response.Name = append(response.Name, values.Name)
			response.Dues = append(response.Dues, values.Dues)
		}
	}

	if len(response.Name) == 0 {
		fmt.Println("all dues are cleared")
		return model.UserDuesResponse{}, nil
	}
	return response, nil
}

// dues amount of particular user service
func UserDues(name string, userList map[string]model.User) (float64, error) {

	user, err := util.GetAValidUserWithGivenName(name, userList)
	if err != nil {
		return 0.0, err
	}
	return user.Dues, nil
}

// list of user who have exhausted all credit limit service
func UsersExhaustedCreditLimit(userList map[string]model.User) (model.UserWithNilCreditLimitResponse, error) {

	var response model.UserWithNilCreditLimitResponse
	for key := range userList {
		user := userList[key]
		if int(user.CreditLimit) == 0 {
			response.Name = append(response.Name, user.Name)
		}
	}
	if len(response.Name) == 0 {
		fmt.Println("no user have reached credit limit")
		return model.UserWithNilCreditLimitResponse{}, nil
	}
	return response, nil
}
