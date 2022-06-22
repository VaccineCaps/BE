package repository

import "golang.org/x/crypto/bcrypt"

// hashPassword implement bycrypt for password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash check bycrypted password
func CheckPasswordHash(password, hashed string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

	return err == nil
}
