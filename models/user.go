package models

import "time"

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name" gorm:"not null"`
	Phone        string `json:"phone" gorm:"unique;not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"-" gorm:"not null"` // - یعنی در JSON نمایش داده نشود
	NationalID   string `json:"national_id" gorm:"unique;not null"`
	Role         string `json:"role" gorm:"default:user"`
	IsVerified   bool   `json:"is_verified" gorm:"default:false"`
	VerifiedCode string `json:"verified_code" gorm:"size:6"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (User) TableName() string {
	return "users"
}
