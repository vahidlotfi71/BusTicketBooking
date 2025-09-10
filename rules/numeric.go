package rules

import (
	"fmt"
	"regexp"
	"strconv"
)

// Numeric بررسی عدد بودن مقدار
func Numeric() Rule {
	return func(value string, fieldName string) (bool, string, error) {
		if matched, _ := regexp.MatchString(`^\d+$`, value); !matched {
			return false, fmt.Sprintf("%s must be a number.", fieldName), nil
		}
		return true, "", nil
	}
}

// Min بررسی حداقل مقدار عددی
func Min(min float64) Rule {
	return func(value string, fieldName string) (bool, string, error) {
		// تبدیل به عدد
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return false, fmt.Sprintf("%s must be a number.", fieldName), nil
		}
		if num < min {
			return false, fmt.Sprintf("%s must be at least %v.", fieldName, min), nil
		}
		return true, "", nil
	}
}
