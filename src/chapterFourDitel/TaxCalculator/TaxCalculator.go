package TaxCalculator

import (
	"fmt"
	"math"
)

func TaxCalculator() {

	fmt.Printf("Welcome To The Tax Calculator\n")

	inputting := 0
	for inputting <= 0 {

		fmt.Println("Enter The Citizen Name: ")
		var citizenName string
		fmt.Scanf("%s", &citizenName)

		fmt.Println("Enter The Citizen Earnings For The Year:  ")
		var citizenWage float64
		fmt.Scanf("%f", &citizenWage)

		taxAmountToPay := sumUpTaxAndEarningRate(citizenWage)

		fmt.Printf("Name: %s\tEarning: %f\t TaxAmount: %f", citizenName, citizenWage, taxAmountToPay)

		fmt.Println("IF YOU WANT TO CONTINUE INPUT 0:  ")
		fmt.Scanf("%d", &inputting)

	}
}

func CalculateTax(names []string, earnings []float64) []float64 {

	var taxToPay []float64
	for _, wage := range earnings {
		taxToPay = append(taxToPay, sumUpTaxAndEarningRate(wage))
	}
	return taxToPay
}

func sumUpTaxAndEarningRate(earning float64) float64 {
	rate := taxRate(earning)

	taxAmountToPay := rate * earning

	return math.Round(taxAmountToPay*100) / 100
}

func taxRate(earning float64) float64 {
	if earning > 30000 {
		return 0.20
	}
	return 0.15
}
