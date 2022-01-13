# A-Simple-Pay-Later-Service
This an application created in golang for simple transaction made between user and merchant or user and application itself.

To run this application :- 
Open terminal- go run main.go

Type the below inputs in the terminal.
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

