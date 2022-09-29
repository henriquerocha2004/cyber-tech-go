package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type SupplierHandler struct {
	supplierRepository entities.SupplierRepository
}

func NewSupplierHandler(supplierRepo entities.SupplierRepository) *SupplierHandler {
	return &SupplierHandler{
		supplierRepository: supplierRepo,
	}
}

func (s *SupplierHandler) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var supplier entities.Supplier
	err := ctx.BodyParser(&supplier)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(supplier)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	err = s.supplierRepository.Create(supplier)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in create supplier: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "supplier created successfully",
	})
}

func (s *SupplierHandler) Update(ctx *fiber.Ctx) error {
	validate := validator.New()
	var supplier entities.Supplier
	id := ctx.Params("id")

	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	err := ctx.BodyParser(&supplier)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(supplier)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	supplier.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = s.supplierRepository.Update(supplier)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in update supplier: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "supplier updated successfully",
	})
}

func (s *SupplierHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	supplierId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = s.supplierRepository.Delete(supplierId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error in delete supplier: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&Output{
		Error:   false,
		Message: "supplier deleted successfully",
	})
}

func (s *SupplierHandler) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	supplierId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	service, err := s.supplierRepository.FindOne(supplierId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error search supplier: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&Output{
		Error: false,
		Data:  service,
	})
}

func (s *SupplierHandler) FindAll(ctx *fiber.Ctx) error {
	supplier, err := s.supplierRepository.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&Output{
			Error:   true,
			Message: "Error search suppliers: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&Output{
		Error: false,
		Data:  supplier,
	})
}
