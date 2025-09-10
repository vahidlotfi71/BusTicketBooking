package rules

import "fmt"

// In بررسی اینکه مقدار در لیست مقادیر مجاز باشد
func In(validValues []string) Rule {
	return func(value, filedName string) (bool, string, error) {
		for _, valid := range validValues {
			if value == valid {
				return true, "", nil
			}
		}
		return false, fmt.Sprintf("%s must be one of the values: %v.", filedName, validValues), nil
	}
}
