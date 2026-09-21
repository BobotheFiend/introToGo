package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter The Length OF the Table")
	var userinput int
	fmt.Scanf("%d", &userinput)

	for i := 1; i <= userinput; i++ {
		if i == 1 {
			fmt.Print("\tN")
		}
		fmt.Printf("\tN%d", i)
		for count := range 5 {

			result := Multiply(count+1, i)
			fmt.Printf("\n\n\t%d\n", result)
		}
	}
}

func Multiply(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
