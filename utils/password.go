package utils

import "golang.org/x/crypto/bcrypt"

// GeneratePassword - Generate bcrypt hash password
func GeneratePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// VerifyPassword - Verify bcrypt hash password
func VerifyPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
