package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type OrderServiceStatusHandler struct {
	orderServiceStatusRepository entities.OrderServiceStatusRepository
}

func NewOrderServiceStatusHandler(orderServiceStatusRepo entities.OrderServiceStatusRepository) *OrderServiceStatusHandler {
	return &OrderServiceStatusHandler{
		orderServiceStatusRepository: orderServiceStatusRepo,
	}
}

func (o *OrderServiceStatusHandler) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var status entities.OrderServiceStatus
	err := ctx.BodyParser(&status)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to parse request")
	}

	err = validate.Struct(status)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to validate request: " + err.Error())
	}

	err = o.orderServiceStatusRepository.Create(status)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in create status",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&Output{
		Error:   false,
		Message: "Status Created successfully",
	})
}

func (o *OrderServiceStatusHandler) Update(ctx *fiber.Ctx) error {
	validate := validator.New()
	var status entities.OrderServiceStatus
	id := ctx.Params("id")

	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	err := ctx.BodyParser(&status)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(status)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	status.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = o.orderServiceStatusRepository.Update(status)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in update status: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "status updated successfully",
	})
}

func (o *OrderServiceStatusHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	statusId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = o.orderServiceStatusRepository.Delete(statusId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in delete status: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "status deleted successfully",
	})
}

func (o *OrderServiceStatusHandler) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	statusId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	status, err := o.orderServiceStatusRepository.FindOne(statusId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in get status: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error: false,
		Data:  status,
	})
}

func (o *OrderServiceStatusHandler) FindAll(ctx *fiber.Ctx) error {
	status, err := o.orderServiceStatusRepository.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in get status: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error: false,
		Data:  status,
	})
}
