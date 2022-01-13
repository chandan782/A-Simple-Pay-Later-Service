package main

import (
	"SIMPLPAY/command"
	"SIMPLPAY/model"
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

func processing(reader bufio.Reader) {
	//userList keep tracks of all the onboarded user along with existing user
	userList := make(map[string]model.User)

	//merchantList keep tracks of all the onboarded user along with existing user
	merchantList := make(map[string]model.Merchant)

	//allTransactionDetails keep tracks of all the transaction that is happening between user and merchant on SIMPLPAY
	allTransactionDetails :=make(map[uuid.UUID]model.Transaction)

	for {
		// reading input from console
		consoleInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		}

		// segregation of console input
		consoleInput = strings.Replace(consoleInput, "\n", "", -1)

		//To exit from SIMPLPAY, type exit
		if consoleInput == "exit" {
			fmt.Println("program terminated!")
			fmt.Println("Thanks For Exploring SIMPLPAY")
			os.Exit(0)
		}

		//handler method to process and segregate the console input
		command.NewCommandHandler(consoleInput, userList, merchantList, allTransactionDetails)
	}
}

func main() {

	//reading input from console
	reader := bufio.NewReader(os.Stdin)

	// input processing
	processing(*reader)

	//To exit from SIMPLPAY type exit
}


/*------------------------------console inputs-----------------------------------------
new user user1 u1@users.com 300
new user user2 u2@users.com 400
new user user3 u3@users.com 500
new merchant m1 m1@merchants.com 0.5%
new merchant m2 m2@merchants.com 1.25%
new merchant m3 m3@merchants.com 1.25%
update merchant m2 1.5%
new txn user2 m1 500
new txn user1 m2 300
new txn user1 m3 10
report users-at-credit-limit
new txn user3 m3 200
new txn user3 m3 300
report users-at-credit-limit
report discount m3
report discount m1
report discount m2
new user user4 u4@users.com 600
new txn user4 m2 500
report discount m2
payback user3 100
report total-dues
----------------------------------------------------------------------------------------*/