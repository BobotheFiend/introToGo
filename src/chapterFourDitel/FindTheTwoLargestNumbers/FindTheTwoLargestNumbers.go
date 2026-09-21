package main

import (
	"fmt"
	"math"
)

func main() {
	largestNumber := math.MinInt
	secondLargestNumber := math.MinInt

	fmt.Println("Input A sequence Of 10 numbers to get the two highest number inputted!")

	count := 1

	for count <= 10 {

		fmt.Printf("Input %d\nEnter A Number:\n ", count)
		var currentInput int
		fmt.Scanf("%d", &currentInput)

		if currentInput > largestNumber {
			secondLargestNumber = largestNumber
			largestNumber = currentInput
		}
		if currentInput > secondLargestNumber && currentInput != largestNumber {
			secondLargestNumber = currentInput
		}

		count++
	}

	fmt.Println("The LargestNumber = ", largestNumber)
	fmt.Println("The Second LargestNumber = ", secondLargestNumber)
}
