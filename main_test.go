package main

import (
	"strconv"
	"testing"
)

func Test_Counterstring(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{
			n:        0,
			expected: "",
		},
		{
			n:        1,
			expected: "*",
		},
		{
			n:        2,
			expected: "2*",
		},

		{
			n:        3,
			expected: "*3*",
		},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			got := Counterstring(tt.n)
			if tt.expected != got {
				t.Errorf("Failed for %v, got %v expected %v", tt.n, got, tt.expected)
			}
		})
	}
}
