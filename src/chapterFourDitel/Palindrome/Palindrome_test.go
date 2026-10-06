package main

import "testing"

func TestPalindromeInputOne(t *testing.T) {

	const input = 12321

	actual := Palindrome(input)

	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPalindromeInputTwo(t *testing.T) {

	const input = 55555

	actual := Palindrome(input)

	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPalindromeInputThree(t *testing.T) {

	const input = 45554

	actual := Palindrome(input)

	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPalindromeInputFour(t *testing.T) {

	const input = 11611

	actual := Palindrome(input)

	expected := true

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPalindromeInputFive(t *testing.T) {

	const input = 12312

	actual := Palindrome(input)

	expected := false

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
