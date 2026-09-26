package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter the Amount Of Numbers You Would Like To Input!!!")
	var input int
	fmt.Scanf("%d", &input)

	largest := math.MinInt
	smallest := math.MaxInt

	for count := 0; count < input; count++ {
		fmt.Printf("Input %d\nEnter A Number!", count+1)
		var userInput int
		fmt.Scanf("%d", &userInput)

		largest = Largest(largest, userInput)
		smallest = Smallest(smallest, userInput)
	}
	sum := Sum(largest, smallest)
	fmt.Printf("The Largest is = %d \nThe Smallest is = %d\n The Sum of %d + %d = %d\n", largest, smallest, largest, smallest, sum)
}

func Largest(largest, number int) int {

	if largest < number {
		largest = number
	}
	return largest
}

func Smallest(smallest, number int) int {
	if number < smallest {
		smallest = number
	}
	return smallest
}

func Sum(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}
