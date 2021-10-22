package jwtmng

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey string
)

func init() {

	signingKey = os.Getenv("SIGNING_KEY")
	if signingKey == "" {
		panic("The environment variable \"SIGNING_KEY\" is not set")
	}
}

func NewJWT(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl * time.Second).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(signingKey))
}

func NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func ValidToken(val string) (jwt.Claims, error) {

	token, err := jwt.Parse(val, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(signingKey), nil
	})

	return token.Claims, err
}

func ParseToken(val string) (jwt.Claims, error) {

	token, err := jwt.ParseWithClaims(val, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	return token.Claims, err
}
