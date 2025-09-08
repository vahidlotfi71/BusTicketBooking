package models

import "time"

// TripType نوع سفر

type TripType string

const (
	TripTypeAirplane TripType = "airplane"
	TripTypeTrain    TripType = "train"
	TripTypeBus      TripType = "bus"
)

// Trip مدل سفر Trip struct{}

type Trip struct {
	ID             uint      `json:"id" ,gorm:"primaryKey"`
	Type           TripType  `json:"type" gorm:"not null"`
	FromCity       string    `json:"from_city" gorm:"not null"`
	ToCity         string    `json:"to_city" gorm:"not null"`
	DepartureTime  time.Time `json:"departure_timeFlightNumber   " gorm:"not null"`
	ArrivalTime    time.Time `json:"arrival_time" gorm:"not null"`
	Price          float64   `json:"price" gorm:"not null"`
	Capacity       uint      `json:"capacity" gorm:"not null"`
	AvailableSeats int       `json:"available_seats" gorm:"not null"`
	Company        string    `json:"company" gorm:"not null"`
	FlightNumber   string    `json:"flight_number" gorm:"not null"` // فقط برای هواپیما
	TrainNumber    string    `json:"train_number" gorm:"not null"`  // فقط برای قطار
	BusNumber      string    `json:"bus_number" gorm:"not null"`    // فقط برای اتوبوس
	IsActive       bool      `json:"is_active" gorm:"default:true"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TableName نام جدول در دیتابیس

func (Trip) TableName() string {
	return "trips"
}
