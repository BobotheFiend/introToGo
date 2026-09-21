package main

import "fmt"

func main() {

	fmt.Print("Enter the number of rows: ")
	var rows int
	fmt.Scanf("%d", &rows)

	stars := "* * * * * * * *"
	for count := 1; count <= rows; count++ {
		if count%2 == 0 {
			fmt.Printf(" %s", stars)
		} else {
			fmt.Printf("%s", stars)
		}
		fmt.Println()
	}

}
