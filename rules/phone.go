package rules

import (
	"fmt"
	"regexp"
)

func phone() Rule {
	phoneRegex := regexp.MustCompile(`^09[0-9]{9}$`)

	return func(value string, fieldName string) (bool, string, error) {
		if !phoneRegex.MatchString(value) {
			return false, fmt.Sprintf("%s must be a valid mobile number.", fieldName), nil
		}
		return true, "", nil
	}
}
