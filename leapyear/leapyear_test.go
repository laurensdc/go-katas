package leapyear

import "testing"

func Test_IsDividableBy4(t *testing.T) {
	if IsLeapYear(4*23) != true {
		t.Errorf("Expected leap year to be dividable by 4")
	}
}

func Test_IsNotDividableBy100(t *testing.T) {
	if IsLeapYear(200) != false {
		t.Errorf("Expected leap year to be not dividable by 100")
	}
}

func Test_IsDividableBy100_WhenAlsoDividableBy400(t *testing.T) {
	if IsLeapYear(800) != true {
		t.Errorf("Expected leap year to be dividable by 100 if also dividable by 400")
	}
}

func Test_NotDividableBy4_IsNotALeapYear(t *testing.T) {
	if IsLeapYear(111) != false {
		t.Errorf("Expected not dividable by 4 not to be a leap year")
	}
}
