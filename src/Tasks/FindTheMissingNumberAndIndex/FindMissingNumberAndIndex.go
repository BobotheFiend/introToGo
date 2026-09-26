package main

import "fmt"

func main() {}

func FindMissingNumberAndIndex(collection []int) []int {

	var result []int

	counter := collection[0]
	index := 0

	for count := 1; count < len(collection); count++ {
		index += 1
		counter += 1
		fmt.Println(counter, collection[count])

		if collection[count] != counter {
			fmt.Println(counter, index)
			result = append(result, counter, index)
			fmt.Println(result)
			counter = collection[count]
			index += 1
		}

	}
	return result
}
