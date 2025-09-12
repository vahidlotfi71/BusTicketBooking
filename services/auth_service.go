package services

import (
	"github.com/vahidlotfi71/BusTicketBooking/config"
	"github.com/vahidlotfi71/BusTicketBooking/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{db: db, cfg: cfg}
}

// CreateUser ایجاد کاربر جدید
func (s *AuthService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

// GetUserByPhone گرفتن کاربر بر اساس شماره موبایل
func (s *AuthService) GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := s.db.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

// GetUserByID گرفتن کاربر بر اساس ID
func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	return &user, err
}

// UpdateUser به‌روزرسانی اطلاعات کاربر
func (s *AuthService) UpdateUser(user *models.User) error {
	return s.db.Save(user).Error
}

// HashPassword هش کردن رمز عبور
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword بررسی صحت رمز عبور
func (s *AuthService) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
