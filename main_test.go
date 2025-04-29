package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	n := 0
	got := Counterstring(n)
	expected := ""
	if got != expected {
		t.Errorf("Failed for %v\ngot %v\nexpected %v\n", n, got, expected)
	}

}

