package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DB        DBConfig
	JWTSecret string
	JWTExpire int
	SMS       SMSConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Charset  string
}

type SMSConfig struct {
	APIKey string
	Sender string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("هشدار: فایل .env یافت نشد، از متغیرهای سیستم استفاده می‌شود")
	}

	return &Config{
		Port: getenv("PORT", "8080"),
		DB: DBConfig{
			Host:     getenv("DB_HOST", "localhost"),
			Port:     getenv("DB_PORT", "3306"),
			User:     getenv("DB_USER", "root"),
			Password: getenv("DB_PASSWORD", ""),
			Name:     getenv("BD_NAME", "ticket_reservation"),
			Charset:  getenv("DB_CHARSET", "utf8mb4"),
		},
		JWTSecret: getenv("JWT_SECRET", "secret-key"),
		JWTExpire: 24,
		SMS: SMSConfig{
			APIKey: getenv("SMS_API_KEY", ""),
			Sender: getenv("SMS_SENDER", ""),
		},
	}

}

func getenv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
