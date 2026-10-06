package ReverseDigits

import "testing"

func TestReverseDigit(t *testing.T) {

	input := 123

	actual := reverse(input)

	expected := true

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}

}

func TestTwoReverseDigit(t *testing.T) {

	input := 1800

	actual := reverse(input)

	expected := false

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}

}

func TestThreeReverseDigit(t *testing.T) {

	input := 526

	actual := reverse(input)

	expected := true

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}

}
