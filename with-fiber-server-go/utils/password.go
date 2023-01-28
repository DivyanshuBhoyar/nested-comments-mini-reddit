package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password cannot be empty")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hashed), err
}


func CheckPassword(savedPass, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(savedPass), []byte(plain))
	return err == nil
}