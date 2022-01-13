package controller

import (
	"SIMPLPAY/model"
	"SIMPLPAY/service"
	"errors"
	"strconv"
	"strings"
)

type UserController interface {
	GetUserWithName(string) (model.User, error)
	UserOnBoard(string, map[string]model.User, []model.User) (model.User, error)
	AllUserWithDues(map[string]model.User) (model.UserDuesResponse, error)
	UserDues(string, map[string]model.User) (float64, error)
	UsersExhaustedCreditLimit(map[string]model.User) (model.UserWithNilCreditLimitResponse, error)
}

// user onboard controller
func UserOnBoard(consoleInput string, userList map[string]model.User) (model.User, error){
	var response model.User
	userData := strings.Split(consoleInput, " ")
	if len(userData) < 5 {
		return model.User{}, errors.New("invalid input, please check and re-enter")
	}

	name := userData[2]
	email := userData[3]
	//input credit limit is a string
	creditLimitStr := userData[4]

	//converting string credit limit to a Float value
	creditLimit, err := strconv.ParseFloat(creditLimitStr, 64)
	if err != nil {
		return response, err
	}

	response, err = service.UserOnBoard(name, email, creditLimit, userList)
	if err != nil {
		return response, err
	}
	return response, nil
}

// list of user whose dues are pending controller
func AllUserWithDues(userList map[string]model.User) (model.UserDuesResponse, error) {
	response, err := service.AllUserWithDues(userList)
	if err != nil {
		return response, err
	}

	return response, nil
}

// dues amount of particular user controller
func UserDues(consoleInput string, userList map[string]model.User) (float64, error) {

	userData := strings.Split(consoleInput, " ")
	if len(userData) < 3 {
		return 0.0, errors.New("invalid input, please check and re-enter")
	}
	name := userData[2]
	response, err := service.UserDues(name, userList)
	if err != nil {
		return 0.0, err
	}

	return response, nil
}

// list of user who have exhausted all credit limit controller
func UsersExhaustedCreditLimit(userList map[string]model.User) (model.UserWithNilCreditLimitResponse, error) {

	response, err := service.UsersExhaustedCreditLimit(userList)
	if err != nil {
		return response, err
	}

	return response,nil
}


