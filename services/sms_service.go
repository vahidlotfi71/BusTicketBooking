package services

import (
	"fmt"

	"github.com/vahidlotfi71/BusTicketBooking/config"
)

type SMSService struct {
	cfg *config.Config
}

func NewSMSService(cfg *config.Config) *SMSService {
	return &SMSService{cfg: cfg}
}

// SendSMS ارسال پیامک (در حالت توسعه فقط لاگ می‌کند)
func (s *SMSService) SendSMS(to string, message string) error {
	// در حالت واقعی اینجا باید به سرویس پیامک متصل شوید
	// برای حالا فقط لاگ می‌کنیم
	fmt.Printf("📨 ارسال پیامک به %s: %s\n", to, message)
	return nil
}
