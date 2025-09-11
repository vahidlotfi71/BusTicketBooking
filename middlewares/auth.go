package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/BusTicketBooking/config"
)

// AuthMiddleware بررسی توکن JWT برای تمام مسیرهای نیازمند لاگین
func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// گرفتن توکن از هدر Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "توکن ارسال نشده است",
			})
		}

		// حذف کلمه Bearer از ابتدای توکن
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// پارس و اعتبارسنجی توکن
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// بررسی روش امضا
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(cfg.JWTSecret), nil
		})

		// اگر توکن نامعتبر بود
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "توکن نامعتبر است",
			})
		}

		// استخراج claims (اطلاعات داخل توکن)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "توکن نامعتبر است",
			})
		}

		// ذخیره اطلاعات کاربر در Locals برای استفاده بعدی
		c.Locals("userID", uint(claims["user_id"].(float64)))
		c.Locals("phone", claims["phone"].(string))
		c.Locals("role", claims["role"].(string))

		// ادامه مسیر
		return c.Next()
	}
}
