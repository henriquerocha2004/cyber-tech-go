package handlers

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/auth"
)

type HandlerAuth struct {
	login auth.Login
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
		log.Println("error in check credentials")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid credentials")
	}
	return ctx.Status(fiber.StatusAccepted).JSON(response)
}
