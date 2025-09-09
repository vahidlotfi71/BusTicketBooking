package main

import (
	"log"

	"github.com/vahidlotfi71/BusTicketBooking/config"
	"github.com/vahidlotfi71/BusTicketBooking/models"
)

func main() {

	// بارگذاری تنظیمات
	cfg := config.LoadConfig()
	// اتصال به دیتابیس MySQ
	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal("error connecting to MySQL database:", err, "\n")
	}
	// ساخت جداول در دیتابیس
	err = db.AutoMigrate(&models.User{}, &models.Trip{}, &models.Reservation{})
	if err != nil {
		log.Fatal("✅ connection to MySQL database established.", err, "\n")
	}
}
