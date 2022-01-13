package constants

const (
	PrefixCommandToCreateNewUser = "new user"
	PrefixCommandToCreateNewMerchant = "new merchant"
	PrefixCommandForTransaction = "new txn"
	PrefixCommandToUpdateMerchant = "update merchant"
	PrefixCommandToPayBack = "payback"
	PrefixCommandToReportDiscount = "report discount"
	PrefixCommandToReportDues = "report dues"
	PrefixCommandToReportUserAtCreditLimit = "report users-at-credit-limit"
	PrefixCommandToReportTotalDues = "report total-dues"
)

const (
	PayBackDestinationNameForUser = "SIMPLPAY"
)

const (
	UserToMerchantTransaction = "User_To_Merchant"
	UserToPayBackTransaction = "User_To_SIMPLPAY"
)