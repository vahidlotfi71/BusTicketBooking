package services

import (
	"time"

	"github.com/vahidlotfi71/BusTicketBooking/models"
	"gorm.io/gorm"
)

type ReservationService struct {
	db *gorm.DB
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{db: db}
}

// CreateReservation ایجاد رزرو جدید
func (s *ReservationService) CreateReservation(reservation *models.Reservation) error {
	return s.db.Create(reservation).Error
}

// GetUserReservations گرفتن رزروهای یک کاربر
func (s *ReservationService) GetUserReservations(userID uint) ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	err := s.db.Preload("Trip").Where("user_id = ?", userID).Find(&reservations).Error
	return reservations, err
}

// GetReservationByID گرفتن رزرو بر اساس ID
func (s *ReservationService) GetReservationByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation
	err := s.db.Preload("Trip").Preload("User").First(&reservation, id).Error
	return &reservation, err
}

// UpdateReservation به‌روزرسانی رزرو
func (s *ReservationService) UpdateReservation(reservation *models.Reservation) error {
	return s.db.Save(reservation).Error
}

// CancelReservation لغو رزرو
func (s *ReservationService) CancelReservation(id uint, userID uint) error {
	// یافتن رزرو
	reservation, err := s.GetReservationByID(id)
	if err != nil {
		return err
	}

	// بررسی اینکه رزرو متعلق به کاربر باشد
	if reservation.UserID != userID {
		return gorm.ErrRecordNotFound
	}

	// بررسی اینکه رزرو قابل لغو باشد
	if reservation.Status != models.StatusPending && reservation.Status != models.StatusConfirmed {
		return gorm.ErrInvalidData
	}

	// تغییر وضعیت و افزودن زمان لغو
	now := time.Now()
	reservation.Status = models.StatusCancelled
	reservation.CancelledAt = &now

	// آزاد کردن صندلی‌ها
	reservation.Trip.AvailableSeats += reservation.SeatsCount
	s.db.Save(&reservation.Trip)

	return s.db.Save(reservation).Error
}
