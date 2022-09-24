package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type ProductCategoryHandler struct {
	productCategoryRepository entities.ProductCategoryRepository
}

func NewProductCategoryHandler(prodCatRepo entities.ProductCategoryRepository) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		productCategoryRepository: prodCatRepo,
	}
}

func (p *ProductCategoryHandler) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var category entities.ProductCategory
	err := ctx.BodyParser(&category)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to parse request")
	}

	err = validate.Struct(validate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("failed to validate request: " + err.Error())
	}

	err = p.productCategoryRepository.Create(category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in create category: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "category created successfully",
	})
}

func (p *ProductCategoryHandler) Update(ctx *fiber.Ctx) error {
	validate := validator.New()
	var category entities.ProductCategory
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	err := ctx.BodyParser(&category)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(category)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	category.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = p.productCategoryRepository.Update(category)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in update category: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "category updated successfully",
	})
}

func (p *ProductCategoryHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	categoryId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = p.productCategoryRepository.Delete(categoryId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in delete category: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "category deleted successfully",
	})
}

func (p *ProductCategoryHandler) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	categoryId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	category, err := p.productCategoryRepository.FindOne(categoryId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error search category: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&Output{
		Error: false,
		Data:  category,
	})
}

func (p *ProductCategoryHandler) FindAll(ctx *fiber.Ctx) error {
	categories, err := p.productCategoryRepository.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error search categories: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&Output{
		Error: false,
		Data:  categories,
	})
}
