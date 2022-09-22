package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/actions"
)

type UserHandler struct {
	userActions actions.UserAction
}

func NewUserHandler(userActions actions.UserAction) *UserHandler {
	return &UserHandler{
		userActions: userActions,
	}
}

func (u *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	validate := validator.New()

	var userInput actions.UserInput
	err := ctx.BodyParser(&userInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(userInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	output := u.userActions.Create(userInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusCreated).JSON(output)
}

func (u *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	validate := validator.New()
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	var userInput actions.UserInput
	err := ctx.BodyParser(&userInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(userInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	userInput.Id, err = strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.Update(userInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.Delete(userId)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) FindAll(ctx *fiber.Ctx) error {
	typeUser := ctx.Query("type_user")
	if typeUser == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("user type not provided")
	}
	output := u.userActions.FindAll(typeUser)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.FindById(userId)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) CreateAddress(ctx *fiber.Ctx) error {
	validate := validator.New()
	var addressInput actions.AddressInput
	err := ctx.BodyParser(&addressInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(addressInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	output := u.userActions.CreateAddress(addressInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusCreated).JSON(output)
}

func (u *UserHandler) UpdateAddress(ctx *fiber.Ctx) error {
	validate := validator.New()
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	var addressInput actions.AddressInput
	err := ctx.BodyParser(&addressInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}
	log.Println(addressInput)
	err = validate.Struct(addressInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	addressInput.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.UpdateAddress(addressInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) DeleteAddress(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	addressId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.DeleteAddress(addressId)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) CreateContact(ctx *fiber.Ctx) error {
	validate := validator.New()
	var contactInput actions.ContactInput
	err := ctx.BodyParser(&contactInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(contactInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	output := u.userActions.CreateContact(contactInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusCreated).JSON(output)
}

func (u *UserHandler) UpdateContact(ctx *fiber.Ctx) error {
	validate := validator.New()
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	var contactInput actions.ContactInput
	err := ctx.BodyParser(&contactInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(contactInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	contactInput.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.UpdateContact(contactInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (u *UserHandler) DeleteContact(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	contactId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := u.userActions.DeleteContact(contactId)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}
