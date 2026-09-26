package main

import "testing"

func TestSumOfSeriesThatRangesTo10(t *testing.T) {

	rangeOfNumbers := 10

	actual := SumSeries(rangeOfNumbers)

	expected := 55

	if actual != expected {
		t.Errorf("SumOfSeriesThatRangesTo10(%d): expected %d, actual %d", rangeOfNumbers, expected, actual)
	}
}

func TestSumOfSeriesThatRangesTo100(t *testing.T) {

	rangeOfNumbers := 100

	actual := SumSeries(rangeOfNumbers)

	expected := 5050

	if actual != expected {
		t.Errorf("SumOfSeriesThatRangesTo10(%d): expected %d, actual %d", rangeOfNumbers, expected, actual)
	}
}

func TestSumOfSeriesThatRangesTo20(t *testing.T) {

	rangeOfNumbers := 20

	actual := SumSeries(rangeOfNumbers)

	expected := 210

	if actual != expected {
		t.Errorf("SumOfSeriesThatRangesTo10(%d): expected %d, actual %d", rangeOfNumbers, expected, actual)
	}
}
