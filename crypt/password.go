package crypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHash 加密
func PasswordHash(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return fmt.Sprintf("%s", b), err
}

// PasswordVerify 校验
func PasswordVerify(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
