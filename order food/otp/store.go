package otp

import (
	"sync"
	"time"
)

type otpEntry struct {
	Code      string
	ExpiresAt time.Time
}

var otpStore = make(map[string]otpEntry)
var mu sync.Mutex

func SetOTP(phone, code string) {
	mu.Lock()
	defer mu.Unlock()
	otpStore[phone] = otpEntry{
		Code:      code,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func VerifyOTP(phone, code string) bool {
	mu.Lock()
	defer mu.Unlock()

	entry, exists := otpStore[phone]
	if !exists || time.Now().After(entry.ExpiresAt) {
		return false
	}
	return entry.Code == code
}
