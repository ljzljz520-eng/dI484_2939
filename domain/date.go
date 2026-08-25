package domain

import "time"

func ValidDate(value string) bool {
	if len(value) != len("2006-01-02") {
		return false
	}
	_, err := time.Parse("2006-01-02", value)
	return err == nil
}

func CompareDate(left, right string) int {
	if left < right {
		return -1
	}
	if left > right {
		return 1
	}
	return 0
}

func DateParts(value string) (int, int, int, bool) {
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return 0, 0, 0, false
	}
	year, month, day := t.Date()
	return year, int(month), day, true
}
