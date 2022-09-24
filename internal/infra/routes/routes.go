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
	api.Post("/service", di.GetServiceHandler().Create)
	api.Put("/service/:id", di.GetServiceHandler().Update)
	api.Delete("/service/:id", di.GetServiceHandler().Delete)
	api.Get("/services", di.GetServiceHandler().FindAll)
	api.Get("service/:id", di.GetServiceHandler().FindOne)
	api.Post("/product-category", di.GetProductCategoryHandler().Create)
	api.Put("/product-category/:id", di.GetProductCategoryHandler().Update)
	api.Delete("/product-category/:id", di.GetProductCategoryHandler().Delete)
	api.Get("/product-categories", di.GetProductCategoryHandler().FindAll)
	api.Get("/product-category/:id", di.GetProductCategoryHandler().FindOne)
	api.Post("/product", di.GetProductHandler().Create)
	api.Put("/product/:id", di.GetProductHandler().Update)
	api.Delete("/product/:id", di.GetProductHandler().Delete)
	api.Get("/products", di.GetProductHandler().FindAll)
	api.Get("/product/:id", di.GetProductHandler().FindOne)
}
