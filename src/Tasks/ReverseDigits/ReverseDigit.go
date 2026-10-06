package ReverseDigits

import (
	"strconv"
	"strings"
)

func reverse(number int) bool {

	return number%10 != 0

}

func reversee(number int) bool {
	var inputToString string
	inputToString = strconv.Itoa(number)

	return !strings.HasSuffix(inputToString, "0")
}
