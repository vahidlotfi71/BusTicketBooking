package rules

import (
	"fmt"
	"regexp"
)

// Email بررسی می‌کند که فرمت ایمیل معتبر باشد
func Email() Rule {
	// عبارت منظم برای بررسی ایمیل
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return func(value string, fieldName string) (bool, string, error) {
		if !emailRegex.MatchString(value) {
			return false, fmt.Sprintf("%s format is not valid.", fieldName), nil
		}
		return true, "", nil
	}
}
