package models

import "time"

type ReservationStatus string

const (
	StatusPending   ReservationStatus = "pending"   // رزرو در انتظار تأیید
	StatusConfirmed ReservationStatus = "confirmed" // رزرو تأیید و قطعی
	StatusCancelled ReservationStatus = "cancelled" //رزرو لغو
	StatusExpired   ReservationStatus = "expired"   //رزرو منقضی شده
)

// Reservation مدل رزرو بلیط

type Reservation struct {
	ID              uint              `json:"id" gorm:"primaryKey"`
	UserID          uint              `json:"user_id" gorm:"not null"`
	TripID          uint              `json:"trip_id" gorm:"not null"`
	Status          ReservationStatus `json:"status" gorm:"default:pending"`
	SeatsCount      int               `json:"seat_count" gorm:"not null"`
	TotalPrice      float64           `json:"total_price" gorm:"not null"`
	SeatNumbers     string            `json:"seat_numbers"` // شماره صندلی‌ها با کاما
	ReservationCode string            `json:"reservation_code" gorm:"unique,not null"`
	CancelledAt     *time.Time        `json:"cancelled_at,omitempty"`
	// دلیل استفاده از پوینتر اینه که ممکنه رزرو لغو نشده باشه.پس مقدارش می‌تونه nil باشه (یعنی لغو نشده).
	// omitempty باعث می‌شه اگر مقدار nil بود، توی JSON فرستاده نشه
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	User User `json:"user" gorm:"foreignKey:UserID"`
	Trip Trip `json:"trip" gorm:"foreignKey:TripID"`
}

func (Reservation) TableName() string {
	return "reservation"
}
