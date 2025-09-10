package rules

import (
	"regexp"
	"strconv"
)

// NationalID بررسی می‌کند که کد ملی معتبر باشد
func NationalID() Rule {
	return func(value string, fieldName string) (bool, string, error) {
		// بررسی فرمت (فقط 10 رقم)
		if matched, _ := regexp.MatchString(`^\d{10}$`, value); !matched {
			return false, "کد ملی باید 10 رقم باشد", nil
		}

		// الگوریتم اعتبارسنجی کد ملی
		sum := 0
		for i := 0; i < 9; i++ {
			digit, _ := strconv.Atoi(string(value[i]))
			sum += digit * (10 - i)
		}

		remainder := sum % 11
		checkDigit, _ := strconv.Atoi(string(value[9]))

		// بررسی رقم کنترل
		if remainder < 2 {
			if checkDigit != remainder {
				return false, "کد ملی معتبر نیست", nil
			}
		} else {
			if checkDigit != (11 - remainder) {
				return false, "کد ملی معتبر نیست", nil
			}
		}

		return true, "", nil
	}
}
