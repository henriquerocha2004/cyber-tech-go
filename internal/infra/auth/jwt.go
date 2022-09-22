package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JwtClaim struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserId    int    `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(email, firstName, lastName string, userId int) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &JwtClaim{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserId:    userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(viper.GetString("auth.jwtKey")))
	return tokenString, err
}

func validateToken(signedToken string) (*JwtClaim, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("auth.jwtKey")), nil
	})
	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var jwtUser JwtClaim
		jwtUser.Email = claim["email"].(string)
		jwtUser.UserId = claim["user_id"].(int)
		jwtUser.FirstName = claim["first_name"].(string)
		jwtUser.LastName = claim["last_name"].(string)
		return &jwtUser, nil
	}

	return nil, errors.New("invalid token or expired")
}
