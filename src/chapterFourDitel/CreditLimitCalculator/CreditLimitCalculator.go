package CreditLimitCalculator

import "fmt"

func CalculateCreditLimit(accountNumber, balanceAtMonthStart, itemsChargedTotal, totalCreditApplied, allowedCreditLimit int) string {
	var eligibility string
	newBalance := (balanceAtMonthStart + itemsChargedTotal) - totalCreditApplied
	fmt.Printf("Your New Balance = %d\nYour Allowed Credit Limit Was = %d\n", newBalance, allowedCreditLimit)

	if newBalance > allowedCreditLimit {
		fmt.Printf("Your New Balance is too big\n")
		eligibility = "Credit Limit HAS Been Exceeded"
	} else {
		eligibility = "Credit Limit has NOT been Exceeded"
	}

	return eligibility
}
