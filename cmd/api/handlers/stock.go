package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/actions"
)

type StockHandler struct {
	stockAction actions.StockActions
}

func NewStockHandler(stockActions actions.StockActions) *StockHandler {
	return &StockHandler{
		stockAction: stockActions,
	}
}

func (s *StockHandler) Add(ctx *fiber.Ctx) error {
	validate := validator.New()
	var stockInput actions.StockInput
	err := ctx.BodyParser(&stockInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(stockInput)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate data: " + err.Error())
	}

	output := s.stockAction.Add(stockInput)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return ctx.Status(fiber.StatusCreated).JSON(output)
}

func (s *StockHandler) FindStock(ctx *fiber.Ctx) error {
	id := ctx.Params("productId")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	output := s.stockAction.FindStock(productId)
	if output.Error {
		return ctx.Status(fiber.StatusInternalServerError).JSON(output)
	}

	return ctx.Status(fiber.StatusCreated).JSON(output)
}
