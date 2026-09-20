package GasMileage

import "testing"

func TestGasMilage(t *testing.T) {

	miles := 300
	gallon := 10

	actual := milesDividedByGallon(miles, gallon)

	/* Assert That */
	expected := 30.0

	if actual != expected {
		t.Errorf("got %f, want %f", actual, expected)
	}

}

func TestGasMilageInputTwo(t *testing.T) {

	miles := 450
	gallon := 12

	actual := milesDividedByGallon(miles, gallon)

	/* Assert That */
	expected := 37.50

	if actual != expected {
		t.Errorf("got %f, want %f", actual, expected)
	}

}

func TestGasMilageInputThree(t *testing.T) {

	miles := 200
	gallon := 8

	actual := milesDividedByGallon(miles, gallon)

	/* Assert That */
	expected := 25.00

	if actual != expected {
		t.Errorf("got %f, want %f", actual, expected)
	}

}

func TestGasMilageInputFour(t *testing.T) {

	miles := 564
	gallon := 22

	actual := milesDividedByGallon(miles, gallon)

	/* Assert That */
	expected := 25.64

	if actual != expected {
		t.Errorf("got %f, want %f", actual, expected)
	}

}
