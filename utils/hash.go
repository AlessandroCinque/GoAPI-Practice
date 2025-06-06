package utils

import "golang.org/x/crypto/bcrypt"

func HashPasssword(passsword string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(passsword), 14)

	return string(bytes), err

}