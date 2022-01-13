package service

import (
	"SIMPLPAY/model"
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserServiceSuite struct {
	userList map[string]model.User
	suite.Suite
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

func (suite *UserServiceSuite) SetupTest() {
	suite.userList = mockUserList()
}

func (suite UserServiceSuite) TestUserOnBoardServiceForSuccessScenario() {

	creditLimit := 300
	actual, _ := UserOnBoard("user2", "u2@users.com", float64(creditLimit), suite.userList)
	expected := model.User{
		Name:          "user2",
		Email:         "u2@users.com",
		CreditLimit:   300.0,
		Dues:          0.0,
		TransactionID: nil,
	}

	suite.Equal(expected, actual)
}

func (suite UserServiceSuite) TestUserOnBoardServiceForErrorScenario() {

	creditLimit := 300
	_, err := UserOnBoard("user1", "u1@users.com", float64(creditLimit), suite.userList)
	expectedErr := errors.New("user already exist with same email id")

	suite.Equal(expectedErr, err)
}

func (suite UserServiceSuite) TestAllUserWithDuesServiceForSuccessScenario() {

	actual,_:=AllUserWithDues(suite.userList)
	expected := model.UserDuesResponse{}
	expected.Name = append(expected.Name,"user3")
	expected.Dues = append(expected.Dues, 100.0)
	expected.Name = append(expected.Name,"user4")
	expected.Dues = append(expected.Dues, 100.0)

	suite.Equal(expected, actual)
}

func (suite UserServiceSuite) TestAllUserWithDuesServiceForErrorScenario() {

	userList := make(map[string]model.User)
	userList["user1"] = model.User{
		Name:          "user1",
		Email:         "u1@users.com",
		CreditLimit:   200.0,
		Dues:          0.0,
		TransactionID: nil,
	}
	_,err:=AllUserWithDues(userList)

	suite.Nil(err)
}

func (suite UserServiceSuite) TestUserDuesServiceForSuccessScenario() {

	actual,_:=UserDues("user3",suite.userList)
	expected := 100.0

	suite.Equal(expected, actual)
}

func (suite UserServiceSuite) TestUserDuesServiceForErrorScenario() {

	_,err:=UserDues("user2",suite.userList)
	expectedErr := errors.New("user with such name doesn't exist")

	suite.Equal(expectedErr, err)
}

func (suite UserServiceSuite) TestUsersExhaustedCreditLimitServiceForSuccessScenario() {

	actual,_:=UsersExhaustedCreditLimit(suite.userList)
	expected := model.UserWithNilCreditLimitResponse{}
	expected.Name = append(expected.Name, "user4")

	suite.Equal(expected, actual)
}

func (suite UserServiceSuite) TestUsersExhaustedCreditLimitServiceForErrorScenario() {

	userList := make(map[string]model.User)
	userList["user1"] = model.User{
		Name:          "user1",
		Email:         "u1@users.com",
		CreditLimit:   0.0,
		Dues:          200.0,
		TransactionID: nil,
	}
	_,err:=UsersExhaustedCreditLimit(suite.userList)
	suite.Nil(err)
}

func mockUserList() map[string]model.User {
	userList := make(map[string]model.User)
	userList["user1"] = model.User{
		Name:          "user1",
		Email:         "u1@users.com",
		CreditLimit:   200.0,
		Dues:          0.0,
		TransactionID: nil,
	}
	userList["user3"] = model.User{
		Name:          "user3",
		Email:         "u1@users.com",
		CreditLimit:   200.0,
		Dues:          100.0,
		TransactionID: nil,
	}
	userList["user4"] = model.User{
		Name:          "user4",
		Email:         "u1@users.com",
		CreditLimit:   0.0,
		Dues:          100.0,
		TransactionID: nil,
	}
	return userList
}
