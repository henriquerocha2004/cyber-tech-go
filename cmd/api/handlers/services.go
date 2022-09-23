package handlers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/entities"
)

type ServiceOutput struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ServiceHandler struct {
	serviceRepository entities.ServiceRepository
}

func NewServiceHandler(serviceRepo entities.ServiceRepository) *ServiceHandler {
	return &ServiceHandler{
		serviceRepository: serviceRepo,
	}
}

func (s *ServiceHandler) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var service entities.Service
	err := ctx.BodyParser(&service)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(service)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	err = s.serviceRepository.Create(service)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&ServiceOutput{
			Error:   true,
			Message: "Error in create service: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&ServiceOutput{
		Error:   false,
		Message: "service created successfully",
	})
}

func (s *ServiceHandler) Update(ctx *fiber.Ctx) error {
	validate := validator.New()
	var service entities.Service
	id := ctx.Params("id")

	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	err := ctx.BodyParser(&service)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = validate.Struct(service)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).SendString("error in validate request")
	}

	service.Id, err = strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = s.serviceRepository.Update(service)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&ServiceOutput{
			Error:   true,
			Message: "Error in update service: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&ServiceOutput{
		Error:   false,
		Message: "service updated successfully",
	})
}

func (s *ServiceHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	serviceId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	err = s.serviceRepository.Delete(serviceId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&ServiceOutput{
			Error:   true,
			Message: "Error in delete service: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(&ServiceOutput{
		Error:   false,
		Message: "service deleted successfully",
	})
}

func (s *ServiceHandler) FindAll(ctx *fiber.Ctx) error {
	services, err := s.serviceRepository.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&ServiceOutput{
			Error:   true,
			Message: "Error search services: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&ServiceOutput{
		Error: false,
		Data:  services,
	})
}

func (s *ServiceHandler) FindOne(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		log.Println("invalid id provided")
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid id provided")
	}

	serviceId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("failed to convert string in int")
		return ctx.Status(fiber.StatusBadRequest).SendString("error in parse request")
	}

	service, err := s.serviceRepository.FindOne(serviceId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&ServiceOutput{
			Error:   true,
			Message: "Error search service: " + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&ServiceOutput{
		Error: false,
		Data:  service,
	})
}
