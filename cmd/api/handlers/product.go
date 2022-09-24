package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type ProductHandler struct {
	productRepository entities.ProductRepository
}

func NewProductHandler(prodRepo entities.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepository: prodRepo,
	}
}

func (p *ProductHandler) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var product entities.Product
	err := ctx.BodyParser(&product)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to parse request")
	}

	err = validate.Struct(product)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to validate request")
	}

	err = p.productRepository.Create(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in create product: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "product created successfully",
	})

}

func (p *ProductHandler) Update(ctx *fiber.Ctx) error {
	validate := validator.New()
	var product entities.Product
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	err := ctx.BodyParser(&product)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(product)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	product.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = p.productRepository.Update(product)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in update product: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "product updated successfully",
	})
}

func (p *ProductHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = p.productRepository.Delete(productId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in delete product: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "product deleted successfully",
	})
}

func (p *ProductHandler) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	product, err := p.productRepository.FindOne(productId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error search product: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&Output{
		Error: false,
		Data:  product,
	})
}

func (p *ProductHandler) FindAll(ctx *fiber.Ctx) error {
	products, err := p.productRepository.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error search products: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&Output{
		Error: false,
		Data:  products,
	})
}
