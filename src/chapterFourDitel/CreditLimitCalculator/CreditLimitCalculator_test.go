package CreditLimitCalculator

import "testing"

func TestToCalculateCreditLimitInputOne(t *testing.T) {
	const (
		accountNumber       = 12333243
		balanceAtMonthStart = 50000
		itemsChargedTotal   = 38000
		totalCreditApplied  = 34560
		allowedCreditLimit  = 30000
	)

	actual := CalculateCreditLimit(accountNumber, balanceAtMonthStart, itemsChargedTotal, totalCreditApplied, allowedCreditLimit)

	/* Assert That */
	expected := "Credit Limit HAS Been Exceeded"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}

}
