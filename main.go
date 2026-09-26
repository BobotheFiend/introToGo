package main

import (
	"fmt"
	"math"
)

func main() {

	me := "BobotheFiend"
	message := "will GO crazy soon!!"

	fmt.Println(me, message, "\n in....")

	for count := 5; count > 0; count-- {
		fmt.Println("\t", count)
	}

	fmt.Println("Gudhwkjskhy5%5trghfefjkjejkfcf......O_o")
	fmt.Println(string(me[6]))

	//--------------------
	var arr [4]int
	arr[0] = 2
	arr[2] = 2
	arr[3] = 6

	fmt.Println(arr)
	//--------------------

	names := [2]string{
		"Cohort",
		"Thirty",
	}
	fmt.Println(names)

	fmt.Println("the slice = ", names[0:1])
	//------------------

	sliceWithArrays := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println("The answer Should be {c,d, e} \nActual =  ", sliceWithArrays[2:5])

	//---------------------
	items := []string{
		"NATIVES",
		"Thirty-Three",
		"Thirty-Two",
		"Thirty-One",
		"Thirty",
		"Twenty-Nine",
		"Twenty-Eight",
		"Twenty-Seven",
	}
	sliceItUp := make([]string, 4, 20)
	copy(sliceItUp, items[1:5])
	fmt.Println(sliceItUp)
	fmt.Println("OKKKKKK//////// ", items[2])
	fmt.Println("Slice It Up Length is = ", len(sliceItUp))
	native := append(sliceItUp, items[len(items)-1], items[len(items)-2])
	fmt.Println(native)
	/* ------------------- */

	facilitators := make(map[string]string)
	facilitators["Kim"] = "Design-Thinking"
	facilitators["Evans"] = "Critical-Thinking & Data-Science"
	facilitators["Chibuzo"] = "Java"
	facilitators["EOB"] = "Java"
	facilitators["SK"] = "Python"
	facilitators["Chinedu"] = "JavaScript & FrontEnd(WebDesign/Html/CSS/React)"
	facilitators["Ebuka"] = "Python"
	facilitators["Majek"] = "Java/ FrontEnd(WebDesign/Html/CSS/React)"
	facilitators["Achalugo"] = "DataBase"
	facilitators["Esther"] = "DataBase"
	facilitators["Femi"] = "Unknown"

	fmt.Println(facilitators["Femi"])

	//*************************************

	//-----------------
	elements := map[string]string{

		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements["Li"])

	/* ----===============================---- */

	chemicalElements := map[string]map[string]string{
		"H": {
			"Name":  "Hydrogen",
			"State": "Gas",
		},
	}
	fmt.Println(chemicalElements["H"]["State"])

	//%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

	anArray := []int{
		1,
		2,
		3,
		10,
		20,
		30,
	}
	const count = 2
	for _, number := range anArray {

		fmt.Println("This is the range stuff = ", number)
		fmt.Printf("This is The Number %d x %d = %d\n\n", number, count, number*count)
	}

	//%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

	/* classwork begins */

	thisArray := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	lowest := math.MaxInt
	for _, number := range thisArray {

		if number < lowest {
			lowest = number
		}

	}
	fmt.Printf("The Lowest Number Is = %d\n", lowest)

	/*   class work end*/
}
