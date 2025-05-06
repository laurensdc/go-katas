package leapyear

func IsLeapYear(i int) bool {
	if i%100 == 0 {
		if i%400 == 0 {
			return true
		} else {
			return false
		}
	}

	if i%4 == 0 {
		return true
	} else {
		return false
	}

}
