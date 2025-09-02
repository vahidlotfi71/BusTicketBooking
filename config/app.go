package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}
}

func Env(key string) string {
	v := os.Getenv(key)
	return v
}
