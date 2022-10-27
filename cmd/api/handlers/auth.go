package handlers

import (
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/auth"
)

type HandlerAuth struct {
	login auth.Login
}

type AuthResponse struct {
	User entities.User
}

func NewHandlerAuth(login auth.Login) *HandlerAuth {
	return &HandlerAuth{
		login: login,
	}
}

func (a *HandlerAuth) Authenticate(ctx *fiber.Ctx) error {
	validate := validator.New()
	var userCredentials auth.UserCredentials
	err := ctx.BodyParser(&userCredentials)
	if err != nil {
		log.Println("error in parse credentials")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid credentials")
	}

	err = validate.Struct(userCredentials)
	if err != nil {
		log.Println("error in validate credentials")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid credentials")
	}

	response, err := a.login.Authenticate(userCredentials)

	if err != nil {
		log.Println("error in check credentials: " + err.Error())
		return ctx.Status(fiber.StatusUnauthorized).SendString("invalid credentials")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    response.Token,
		HTTPOnly: true,
	})

	var authResponse AuthResponse
	authResponse.User = response.User

	return ctx.Status(fiber.StatusAccepted).JSON(authResponse)
}

func (a *HandlerAuth) CheckMe(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusAccepted).SendString("ok")
}

func (a *HandlerAuth) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("token")
	return ctx.Status(fiber.StatusOK).SendString("logout successfully")
}
