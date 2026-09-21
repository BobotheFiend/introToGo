package main

import "testing"

func TestComparatorInputOne(t *testing.T) {

	const (
		a = 30
		b = 29
	)

	actual := Comparator(a, b)

	expected := "+1"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestComparatorInputTwo(t *testing.T) {

	const (
		a = 70
		b = 70
	)

	actual := Comparator(a, b)

	expected := "0"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestComparatorInputThree(t *testing.T) {

	const (
		a = 277
		b = 290
	)

	actual := Comparator(a, b)

	expected := "-1"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
