package SalesCommisionCalculator

import "testing"

func TestSalesCommissionCalculatorCaseOne(t *testing.T) {

	weeklyEarnings := []float64{239.99, 129.75, 99.95, 350.89}

	actual := CalculateCommission(weeklyEarnings)

	expected := 1094.43

	if actual != expected {
		t.Errorf("got %f want %f", actual, expected)
	}

}
