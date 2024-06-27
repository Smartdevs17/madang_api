package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOTP generates a 4-digit OTP
func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(10000)         // generates a number between 0 and 9999
	return fmt.Sprintf("%04d", otp) // format the number as a 4-digit string
}
