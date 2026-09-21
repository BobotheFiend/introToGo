package main

import "fmt"

func main() {

	fmt.Println("Enter A Huge Number!")
	var initalNumber int
	fmt.Scanf("%d", &initalNumber)

	result := 0

	count := 1
	for count > 0 {

		fmt.Println("Enter A Number: ")
		var userInput int
		fmt.Scanf("%d", &userInput)

		result += userInput

		if result >= initalNumber {
			count = 0
		}
	}

	fmt.Printf("Huge Number: %d\nSum Of Input: %d\n", initalNumber, result)
	fmt.Println("END!")
}
