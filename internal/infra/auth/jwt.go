package auth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
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
	expirationTime := jwt.NewNumericDate(time.Now().Add(10 * time.Hour))
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

func CheckAuth(ctx *fiber.Ctx) error {
	token := ctx.Cookies("token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	err := validateToken(token)
	if err != nil {
		ctx.ClearCookie("token")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Next()
}

func validateToken(signedToken string) error {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("auth.jwtKey")), nil
	})
	if err != nil {
		log.Println(err)
		return err
	}

	if !token.Valid {
		return errors.New("token invalid")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token or expired")
}
