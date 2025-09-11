package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// RequireRole بررسی نقش خاص (مثلاً admin)
func RequireRole(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// گرفتن نقش کاربر از Locals (در middleware قبلی ذخیره شده)
		userRole := c.Locals("role").(string)

		// اگر نقش کاربر با نقش مورد نظر برابر نبود
		if userRole != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "شما اجازه دسترسی به این بخش را ندارید",
			})
		}

		// اگر نقش درست بود، ادامه مسیر
		return c.Next()
	}
}
