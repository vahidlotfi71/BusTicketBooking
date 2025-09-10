package rules

import (
	"fmt"
	"time"
)

// DateTime بررسی فرمت تاریخ و زمان
func DateTime() Rule {
	return func(value string, fieldName string) (bool, strin, error) {
		// تلاش برای پارس کردن تاریخ با فرمت خاص
		_, err := time.Parse("2006-01-02 15:04", value)
		if err != nil {
			return false, fmt.Sprintf("%s must be in the correct date and time format (example: 2024-01-15 14:30).", fieldName), nil
		}
		return true, "", nil
	}
}

// After بررسی اینکه زمان بعد از زمان دیگری باشد
func After() Rule {
	return func(value string, fieldName string) (bool, string, error) {
		// اینجا فقط فرمت را چک می‌کنیم - منطق مقایسه در کنترلر انجام می‌شود
		_, err := time.Parse("2006-01-02 15:04", value)
		if err != nil {
			return false, fmt.Sprintf("%s must be a valid date and time.", fieldName), nil
		}
		return true, "", nil
	}
}
