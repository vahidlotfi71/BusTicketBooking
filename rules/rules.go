package rules

// Rule تابعی که یک قانون اعتبارسنجی را تعریف می‌کند
// ورودی: مقدار فیلد، نام فیلد
// خروجی: آیا قانون رعایت شده، پیام خطا، خطای احتمالی
type Rule func(value string, filedName string) (passed bool, message string, err error)

// FieldRules تعریف قوانین برای هر فیلد
type FieldRules struct {
	Field string `json:"field"` // نام فیلد
	Rules []Rule `json:"rules"` // لیست قوانین
}

// ValidationError خطای اعتبارسنجی
type ValidationError struct {
	Field   string `json:"field"`   // نام فیلد خطادار
	Message string `json:"message"` // پیام خطا
}
