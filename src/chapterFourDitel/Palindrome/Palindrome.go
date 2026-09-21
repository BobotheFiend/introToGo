package main

import "fmt"

func main() {
	fmt.Println("Five Digit Palindrome Checker!!")

	var input int
	count := 0
	for count < 0 {
		fmt.Printf("Enter a five digit number: ")
		fmt.Scanf("%d", &input)

		toString := fmt.Sprintf("%d", input)

		if len(toString) != 5 {
			fmt.Println("Invalid input")
		} else {
			count++
		}
	}

	isPalindrom := Palindrome(input)
	if isPalindrom {
		fmt.Printf("%d is A Palindrome!", input)
	} else {
		fmt.Printf("%d is NOT a Palindrome!", input)
	}
}

func Palindrome(number int) bool {

	convertToString := fmt.Sprintf("%d", number)
	if convertToString[0] == convertToString[4] && convertToString[1] == convertToString[3] {
		return true
	}
	return false
}
