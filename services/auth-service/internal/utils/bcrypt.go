package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(plainPasswordInput string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPasswordInput))
}

func HashPassword(plainPasswordInput string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainPasswordInput), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}