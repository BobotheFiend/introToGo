package main

import "strings"

import "fmt"

func main() {

	fmt.Println("Enter the Length Of Base FOr The Right Angled Triangle To Be Drawn ")
	var userInput int
	fmt.Scanf("%d", &userInput)
	var star strings.Builder

	for i := 1; i <= userInput; i++ {
		star.WriteString("*")
		fmt.Printf("%s \n", star.String())
	}

}
