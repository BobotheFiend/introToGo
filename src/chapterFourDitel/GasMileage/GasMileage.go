package GasMileage

import (
	"fmt"
	"math"
)

func CalculateMilesPerGallon(tripAmount int) float64 {
	var result float64
	counter := 1

	for tripAmount >= counter {

		fmt.Println("Enter a The Miles Driven!: ")
		var milesDriven int
		fmt.Scanf("%d", &milesDriven)

		fmt.Println("Enter a The Gallon Amount Consumed!: ")
		var gallonUsed int
		fmt.Scanf("%d", &gallonUsed)

		sum := milesDividedByGallon(milesDriven, gallonUsed)
		result += sum

		fmt.Printf("Trip %d Cost: %fm/g\n", counter, sum)

		counter++

	}

	fmt.Printf("Your Total gas MileAge From %d Trips....\n Total = %f", tripAmount, result)

	return result
}

func milesDividedByGallon(miles int, gallon int) float64 {
	var result float64
	result = float64(float64(miles) / float64(gallon))
	return math.Round(result*100) / 100
}
