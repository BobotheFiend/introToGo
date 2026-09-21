package main

import "testing"

func TestFindingTheLargestInputOne(t *testing.T) {

	const (
		largest       int = 60
		currentNumber int = 46
	)

	actual := Largest(largest, currentNumber)

	expect := 60
	if actual != expect {
		t.Errorf("Largest = %d, want %d", actual, expect)
	}
}

func TestFindingTheLargestInputTwo(t *testing.T) {

	const (
		largest       int = 38
		currentNumber int = 89
	)

	actual := Largest(largest, currentNumber)

	expect := 89
	if actual != expect {
		t.Errorf("Largest = %d, want %d", actual, expect)
	}
}
