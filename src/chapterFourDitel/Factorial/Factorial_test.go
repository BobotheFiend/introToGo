package main

import "testing"

func TestFactorialForInput5(t *testing.T) {

	input := 5

	actual := Factorial(input)

	expected := 120
	if actual != expected {
		t.Errorf("For input = %d, expected %d, got %d", input, expected, actual)

	}
}

func TestFactorialForInput10(t *testing.T) {

	input := 10

	actual := Factorial(input)

	expected := 3628800
	if actual != expected {
		t.Errorf("For input = %d, expected %d, got %d", input, expected, actual)
	}

}
