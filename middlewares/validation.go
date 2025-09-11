package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vahidlotfi71/BusTicketBooking/rules"
)

// ValidationMiddleware میان‌افزاری برای اعتبارسنجی داده‌ها
func ValidationMiddleware(fieldRules []rules.FieldRules) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// map برای نگهداری داده‌های بدنه درخواست
		body := make(map[string]interface{})

		// تبدیل JSON بدنه به map
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Data format is incorrect.",
			})
		}
		if len(body) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "No data was submitted",
			})
		}
		// لیست خطاها
		var validationErrors []rules.ValidationError

		// بررسی هر فیلد
		for _, fieldRule := range fieldRules {
			// گرفتن مقدار فیلد
			value := ""
			if v, ok := body[fieldRule.Field].(string); ok {
				value = v
			}
			// اعمال قوانین برای این فیلد
			for _, rule := range fieldRule.Rules {
				passed, message, err := rule(value, fieldRule.Field)

				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"success": false,
						"message": "Sever error",
					})
				}
				// اگر قانون نقض شد
				if !passed {
					validationErrors = append(validationErrors, rules.ValidationError{
						Field:   fieldRule.Field,
						Message: message,
					})
					break // از بقیه قوانین این فیلد صرف نظر می‌کنیم
				}
			}
		}
		// اگر خطا وجود داشت
		if len(validationErrors) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"errors":  validationErrors,
			})
		}
		// اگر همه چیز درست بود، برو به کنترلر
		return c.Next()
	}
}
