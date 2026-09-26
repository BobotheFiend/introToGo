package main

import (
	"reflect"
	"testing"
)

func TestMissingNumberAndIndexInputOne(t *testing.T) {

	input := []int{1, 2, 4, 5, 6}

	actual := FindMissingNumberAndIndex(input)

	expected := []int{3, 2}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected", expected, "Got", actual)
	}
}

func TestMissingNumberAndIndexInputTwo(t *testing.T) {

	input := []int{4, 5, 6, 8, 10}

	actual := FindMissingNumberAndIndex(input)

	expected := []int{7, 3, 9, 5}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected", expected, "Got", actual)
	}
}
