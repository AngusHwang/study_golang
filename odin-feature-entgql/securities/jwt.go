package securities

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"odin/ent"
	"os"
	"time"
)

type JwtCustomClaim struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

var keyFunc = func(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("unexpected signing method")
	}

	return jwtSecret, nil
}

func GenerateWithUserTokens(user *ent.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Parse(tokenString string) (*JwtCustomClaim, error) {
	var claims JwtCustomClaim
	_, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
