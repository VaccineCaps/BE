package hash

import "golang.org/x/crypto/bcrypt"

// hashPassword implement bycrypt for password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
