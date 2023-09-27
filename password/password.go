package password

import (
	"github.com/lights-T/goutils/xstrings"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(xstrings.StringToBytes(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(xstrings.StringToBytes(hash), xstrings.StringToBytes(password))
	return err == nil
}
