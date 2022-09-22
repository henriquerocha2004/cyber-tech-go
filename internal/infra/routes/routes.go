package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/di"
)

func Register(app *fiber.App) {
	var di = di.DependencyContainer{}
	api := app.Group("/api")
	api.Post("/auth", di.GetAuthHandler().Authenticate)
	api.Post("/user", di.GetUserHandler().CreateUser)
	api.Put("/user/:id", di.GetUserHandler().UpdateUser)
	api.Delete("/user/:id", di.GetUserHandler().DeleteUser)
	api.Get("/users", di.GetUserHandler().FindAll)
	api.Get("/user/:id", di.GetUserHandler().FindById)
	api.Post("/address", di.GetUserHandler().CreateAddress)
	api.Put("/address/:id", di.GetUserHandler().UpdateAddress)
	api.Delete("/address/:id", di.GetUserHandler().DeleteAddress)
	api.Post("/contact", di.GetUserHandler().CreateContact)
	api.Put("/contact/:id", di.GetUserHandler().UpdateContact)
	api.Delete("/contact/:id", di.GetUserHandler().DeleteContact)
}
