package services

import (
	"github.com/vahidlotfi71/BusTicketBooking/models"
	"gorm.io/gorm"
)

type TripService struct {
	db *gorm.DB
}

func NewTripService(db *gorm.DB) *TripService {
	return &TripService{db: db}
}

// CreateTrip ایجاد سفر جدید
func (s *TripService) CreateTrip(trip *models.Trip) error {
	trip.AvailableSeats = int(trip.Capacity) // در ابتدا همه صندلی‌ها خالی‌اند
	return s.db.Create(trip).Error
}

// GetTripByID گرفتن سفر بر اساس ID
func (s *TripService) GetTripByID(id uint) (*models.Trip, error) {
	var trip models.Trip
	err := s.db.First(&trip, id).Error
	return &trip, err
}

// GetTripsWithFilter گرفتن لیست سفرها با فیلتر و صفحه‌بندی
func (s *TripService) GetTripsWithFilter(fromCity, toCity, tripType, date string, page, limit int) ([]*models.Trip, int64, error) {
	var trips []*models.Trip
	var total int64

	query := s.db.Model(&models.Trip{}).Where("is_active = ?", true)

	if fromCity != "" {
		query = query.Where("from_city LIKE ?", "%"+fromCity+"%")
	}
	if toCity != "" {
		query = query.Where("to_city LIKE ?", "%"+toCity+"%")
	}
	if tripType != "" {
		query = query.Where("type = ?", tripType)
	}
	if date != "" {
		query = query.Where("DATE(departure_time) = ?", date)
	}

	// شمارش کل رکوردها
	query.Count(&total)

	// صفحه‌بندی
	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&trips).Error

	return trips, total, err
}

// UpdateTrip به‌روزرسانی سفر
func (s *TripService) UpdateTrip(trip *models.Trip) error {
	return s.db.Save(trip).Error
}

// DeleteTrip حذف سفر
func (s *TripService) DeleteTrip(id uint) error {
	return s.db.Delete(&models.Trip{}, id).Error
}

// HasActiveReservations بررسی وجود رزرو فعال برای سفر
func (s *TripService) HasActiveReservations(tripID uint) (bool, error) {
	var count int64
	err := s.db.Model(&models.Reservation{}).
		Where("trip_id = ? AND status IN (?)", tripID, []models.ReservationStatus{models.StatusPending, models.StatusConfirmed}).
		Count(&count).Error
	return count > 0, err
}
