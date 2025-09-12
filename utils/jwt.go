package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vahidlotfi71/BusTicketBooking/config"
)

// GenerateJWT تولید توکن JWT
func GenerateJWT(userID uint, phone string, role string, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"phone":   phone,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 ساعت اعتبار
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
