package main

import (
	"testing"
)

func Test_Zero(t *testing.T) {
	n := 0
	got := Counterstring(n)
	expected := ""

	if got != expected {
		t.Errorf("Failed for %v\ngot %v\nexpected %v\n", n, got, expected)
	}
}

func Test_One(t *testing.T) {
	n := 1
	got := Counterstring(n)
	expected := "*"

	if got != expected {
		t.Errorf("Failed for %v\ngot %v\nexpected %v\n", n, got, expected)
	}
}

func Test_Two(t *testing.T) {
	n := 2
	got := Counterstring(n)
	expected := "2*"

	if got != expected {
		t.Errorf("Failed for %v\ngot %v\nexpected %v\n", n, got, expected)
	}
}
