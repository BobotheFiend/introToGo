package SalesCommisionCalculator

import "math"

func CalculateCommission(earnings []float64) float64 {

	var grossSales float64
	const commision int = 200
	const grossPercentage float64 = 0.09

	count := 0

	for count < len(earnings) {
		grossSales += earnings[count]
		count++
	}

	grossSales += float64(commision) + (grossPercentage * grossSales)
	return math.Round(grossSales*100) / 100
}
