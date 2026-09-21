package main

func main() {

}

func Comparator(firstNumber, secondNumber int) string {
	var result string
	if firstNumber == secondNumber {
		result = "0"
	}
	if firstNumber < secondNumber {
		result = "-1"
	}
	if firstNumber > secondNumber {
		result = "+1"
	}
	return result
}
