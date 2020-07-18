package utils

import (
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
)

// GetToken ...
func GetToken(secret string) string {
	token, _ := totp.GenerateCode(secret, time.Now())
	return token
}

// PrintToken ...
func PrintToken(provider string, token string) {
	fmt.Printf("%s: %s\n", provider, token)
}
