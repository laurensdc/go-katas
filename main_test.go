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
		{
			n:        4,
			expected: "2*4*",
		},
		{
			n:        5,
			expected: "*3*5*",
		},
		{
			n:        6,
			expected: "2*4*6*",
		},
		{
			n:        9,
			expected: "*3*5*7*9*",
		},
		{
			n:        10,
			expected: "*3*5*7*10*",
		},
		{
			n:        11,
			expected: "2*4*6*8*11*",
		},
		{
			n:        99,
			expected: "*3*5*7*9*12*15*18*21*24*27*30*33*36*39*42*45*48*51*54*57*60*63*66*69*72*75*78*81*84*87*90*93*96*99*",
		},
		{
			n:        100,
			expected: "*3*5*7*9*12*15*18*21*24*27*30*33*36*39*42*45*48*51*54*57*60*63*66*69*72*75*78*81*84*87*90*93*96*100*",
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
