package util

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GenerateSignedString(userId int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(userId)),
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(userId)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	return token.SignedString(getJWTSecret())
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&MyJWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return getJWTSecret(), nil
		},
	)

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			// token invalid
			err = errors.New("signature validation failed")
			return
		case jwt.ValidationErrorExpired:
			// token expired
			err = errors.New("token is expired")
			return
		default:
			err = errors.New("token is invalid")
			return
		}
	}

	if !token.Valid {
		err = errors.New("unauthorized")
		return
	}

	return
}
