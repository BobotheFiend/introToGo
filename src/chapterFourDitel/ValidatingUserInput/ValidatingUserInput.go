package main

import (
	"fmt"
)

func main() {

	var currentInput int

	count := 0
	for count <= 0 {

		fmt.Println("Enter A Number: ")
		fmt.Scanf("%d", &currentInput)

		fmt.Println(currentInput)
		if currentInput == 2 || currentInput == 1 {
			count += 1
		}
	}

}
