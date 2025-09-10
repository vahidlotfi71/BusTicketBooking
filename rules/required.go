package rules

import "fmt"

// Required بررسی می‌کند که فیلد خالی نباشد

func Required() Rule {
	return func(value string, fieldName string) (bool, string, error) {
		if fieldName == "" {
			return false, fmt.Sprintf("%s field is required"), nil
		}
		return true, "", nil
	}
}
