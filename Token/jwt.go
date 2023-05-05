package Token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"test-dikurium/graph/model"
)

type JWTMaker struct {
	SecretKey string
}

type customClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func NewJWTMaker(secretKey string) JWTMaker {
	return JWTMaker{SecretKey: secretKey}
}

func (J *JWTMaker) CreateToken(user model.User) (string, error) {

	claims := customClaims{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "domain_url",
			Subject:   "Oreoluwa",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(J.SecretKey))
	if err != nil {
		return "", err
	}
	return signedString, err
}

func (J *JWTMaker) VerifyToken(token string) (*customClaims, error) {
	Keyfunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(J.SecretKey), nil
	}

	TokenClaims, err := jwt.ParseWithClaims(token, &customClaims{}, Keyfunc)
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New(" invalid signature error")
		}

		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New(" token expired !")
		}

		return nil, err
	}

	if claims, ok := TokenClaims.Claims.(*customClaims); ok && TokenClaims.Valid {
		return claims, nil
	}
	return nil, errors.New("error converting")

}
