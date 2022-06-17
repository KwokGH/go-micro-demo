package utils

import "golang.org/x/crypto/bcrypt"

// GenerateFromPassword 加密字符串
func GenerateFromPassword(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ComparePassword(hash string, source string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(source)); err != nil {
		return false, err
	}
	return true, nil
}
