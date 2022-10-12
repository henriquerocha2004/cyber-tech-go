package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/actions"
)

type OrderServiceHandler struct {
	orderActions actions.ServiceOrderActions
}

func NewOrderServiceHandler(orderActions actions.ServiceOrderActions) *OrderServiceHandler {
	return &OrderServiceHandler{
		orderActions: orderActions,
	}
}

func (o *OrderServiceHandler) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var orderInput actions.ServiceOrderInput
	err := ctx.BodyParser(&orderInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to parse params")
	}
	err = validate.Struct(orderInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to validate params")
	}

	output := o.orderActions.Create(orderInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusCreated).JSON(output)
}

func (o *OrderServiceHandler) Update(ctx *fiber.Ctx) error {
	validate := validator.New()
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	var orderInput actions.ServiceOrderInput
	err := ctx.BodyParser(&orderInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to parse params")
	}

	err = validate.Struct(orderInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	orderInput.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := o.orderActions.Update(orderInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (o *OrderServiceHandler) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := o.orderActions.GetOne(orderId)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

func (o *OrderServiceHandler) FindAll(ctx *fiber.Ctx) error {
	output := o.orderActions.GetAll()
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

type OrderServiceQueueHandler struct {
	stockAction actions.StockActions
}

func NewOrderServiceQueueHandler(stk actions.StockActions) *OrderServiceQueueHandler {
	return &OrderServiceQueueHandler{
		stockAction: stk,
	}
}

func (o *OrderServiceQueueHandler) Distribute(message string) error {
	var orderInput actions.ServiceOrderInput
	err := json.Unmarshal([]byte(message), &orderInput)
	if err != nil {
		log.Println(err)
		return err
	}

}
