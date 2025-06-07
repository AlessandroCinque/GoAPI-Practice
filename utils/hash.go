package utils

import "golang.org/x/crypto/bcrypt"

func HashPasssword(passsword string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(passsword), 14)

	return string(bytes), err

}

func CHeckPassHash(password, hashPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))

	if err != nil { return false } else { return true }
}