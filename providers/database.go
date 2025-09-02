package providers

import (
	"fmt"
	"log"

	"github.com/vahidlotfi71/BusTicketBooking/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	config.Init()

	dbHost := config.Env("DB_HOST")
	dbPort := config.Env("DB_PORT")
	dbUser := config.Env("DB_USER")
	dbPass := config.Env("DB_PASS")
	dbName := config.Env("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL connection failed: %v", err)
	}
	log.Print("MySQL connected")

}
