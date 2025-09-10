package rules

import (
	"fmt"
)

// MinLength بررسی حداقل طول رشته
func MinLength(min int) Rule {
	return func(value string, fieldName string) (bool, string, error) {
		if len(value) < min {
			return false, fmt.Sprintf("%s must be at least %d characters.", fieldName, min), nil
		}
		return true, "", nil
	}
}

// MaxLength بررسی حداکثر طول رشته
func MaxLength(max int) Rule {
	return func(value string, fieldName string) (bool, string, error) {
		if len(value) > max {
			return false, fmt.Sprintf("%s must be at most %d characters.", fieldName, max), nil
		}
		return true, "", nil
	}
}

// LengthBetween بررسی محدوده طول رشته
func LengthBetween(min int, max int) Rule {
	return func(value, fieldName string) (bool, string, error) {
		lenghth := len(value)
		if lenghth < min || lenghth > max {
			return false, fmt.Sprintf("%s must be between %d and %d characters.", fieldName, min, max), nil
		}
		return true, "", nil
	}
}
