package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashPassword), nil

}

func ComparePassword(hashedPass, enteredPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(enteredPass))

}