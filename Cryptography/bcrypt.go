package Cryptography

import "golang.org/x/crypto/bcrypt"

type BCryptMaker struct {
}

func (c BCryptMaker) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c BCryptMaker) ComparePassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func NewCryptoSrv() BCryptMaker {
	return BCryptMaker{}
}
