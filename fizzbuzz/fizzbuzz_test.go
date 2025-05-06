package fizzbuzz

import (
	"testing"
)

func Test_Returns1To100(t *testing.T) {
	got := len(GenerateNumbers(100))
	expected := 100

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}

func Test_FirstItemIs1(t *testing.T) {
	got := GenerateNumbers(100)
	expected := "1"

	if got[0] != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}

func Test_FirstSecondItemIs2(t *testing.T) {
	got := GenerateNumbers(100)[1]
	expected := "2"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}

func Test_22thItemIs22(t *testing.T) {
	got := GenerateNumbers(100)[21]
	expected := "22"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}

func Test_ItemsDivisibleBy3AreFizz(t *testing.T) {
	got := GenerateNumbers(100)[2]
	expected := "fizz"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}

	got = GenerateNumbers(100)[5]
	expected = "fizz"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}

func Test_ItemsDivisibleBy5AreBuzz(t *testing.T) {
	got := GenerateNumbers(100)[4]
	expected := "buzz"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}

	got = GenerateNumbers(100)[9]
	expected = "buzz"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}

func Test_ItemsDivisible_By_Both3And5_AreGenerateNumbers(t *testing.T) {
	got := GenerateNumbers(100)[14]
	expected := "fizzbuzz"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}

	got = GenerateNumbers(100)[29]
	expected = "fizzbuzz"

	if got != expected {
		t.Errorf("Got %v expected %v\n", got, expected)
	}
}
