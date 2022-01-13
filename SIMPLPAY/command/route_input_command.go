package command

import (
	"SIMPLPAY/constants"
	"SIMPLPAY/controller"
	"SIMPLPAY/model"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func NewCommandHandler(consoleInput string, userList map[string]model.User, merchantList map[string]model.Merchant, allTransactionDetails map[uuid.UUID]model.Transaction)() {

	//command to create new user and added it top existing userList map
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToCreateNewUser) {
		user,err := controller.UserOnBoard(consoleInput, userList)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(user.Name)
		fmt.Print("(")
		fmt.Print(user.CreditLimit)
		fmt.Print(")")
		fmt.Print("\n")
	}

	//command to create new merchant and added it top existing merchantList map
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToCreateNewMerchant) {
		merchant, err := controller.MerchantOnBoard(consoleInput, merchantList)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(merchant.Name)
		fmt.Print("(")
		fmt.Print(merchant.DiscountOffered)
		fmt.Print(")")
		fmt.Print("\n")
	}

	//command for transaction between user and merchant at SIMPLPAY
	if strings.HasPrefix(consoleInput,constants.PrefixCommandForTransaction) {
		transaction, err := controller.Transaction(consoleInput, userList, merchantList, allTransactionDetails)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(transaction.Status)
		fmt.Println("!")
	}

	//command to update merchant discount
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToUpdateMerchant) {
		_, err := controller.UpdateMerchant(consoleInput, merchantList)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("updated!")
	}

	//command for transaction between user and SIMPLPAY
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToPayBack) {
		response, err := controller.UserPayBack(consoleInput, userList, allTransactionDetails)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(response.UserName)
		fmt.Print("(")
		fmt.Print("Dues:", response.Dues)
		fmt.Println(")")
	}

	//command to report merchant overall discounted amount on different transaction
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToReportDiscount) {
		response, err := controller.MerchantDiscountReport(consoleInput, merchantList, allTransactionDetails)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response.TotalDiscount)
	}

	//command to report user dues
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToReportDues) {
		dues, err := controller.UserDues(consoleInput, userList)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dues)
	}

	//command to report users who have reached there credit-limit
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToReportUserAtCreditLimit) {
		response, err := controller.UsersExhaustedCreditLimit(userList)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(response.Name) == 0 {
			return
		}

		for i := range response.Name {
			fmt.Println(response.Name[i])
		}
	}

	//command to report all user who have pending dues
	if strings.HasPrefix(consoleInput,constants.PrefixCommandToReportTotalDues) {
		response, err := controller.AllUserWithDues(userList)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(response.Name) == 0 {
			return
		}

		var totalDues float64
		for i := range response.Name {
			fmt.Print(response.Name[i])
			fmt.Println(":",response.Dues[i])
			totalDues += response.Dues[i]
		}
		fmt.Println("total:",totalDues)
	}
	return
}

