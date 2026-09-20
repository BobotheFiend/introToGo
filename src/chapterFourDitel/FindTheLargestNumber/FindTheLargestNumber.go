package main

import (
	"fmt"
	"math"
)

func main() {

	largestNumber := math.MinInt

	fmt.Println("Input A sequence Of 10 numbers to get thhe highest number inputted!")

	count := 1

	for count <= 10 {

		fmt.Printf("Input %d\nEnter A Number:\n ", count)
		var currentInput int
		fmt.Scanf("%d", &currentInput)

		largestNumber = Largest(largestNumber, currentInput)

		count++
	}

	fmt.Println("The LargestNumber = ", largestNumber)
}

func Largest(currentLargestNumber, target int) int {

	if target > currentLargestNumber {
		currentLargestNumber = target
	}
	return currentLargestNumber
}
