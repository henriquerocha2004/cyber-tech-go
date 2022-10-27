package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/auth"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/di"
)

func Register(app *fiber.App) {
	var di = di.DependencyContainer{}
	api := app.Group("/api")
	api.Post("/auth", di.GetAuthHandler().Authenticate)
	api.Get("/logout", di.GetAuthHandler().Logout, auth.CheckAuth)

	checkMe := api.Group("check", auth.CheckAuth)
	checkMe.Get("/me", di.GetAuthHandler().CheckMe)

	productCategory := api.Group("product-category")
	productCategory.Post("/create", di.GetProductCategoryHandler().Create)
	productCategory.Put("/:id", di.GetProductCategoryHandler().Update)
	productCategory.Delete("/:id", di.GetProductCategoryHandler().Delete)
	productCategory.Get("/", di.GetProductCategoryHandler().FindAll)
	productCategory.Get("/:id", di.GetProductCategoryHandler().FindOne)

	product := api.Group("product", auth.CheckAuth)
	product.Get("/", di.GetProductHandler().FindAll)
	product.Get("/:id", di.GetProductHandler().FindOne)
	product.Put("/:id", di.GetProductHandler().Update)
	product.Delete("/:id", di.GetProductHandler().Delete)
	product.Post("/create", di.GetProductHandler().Create)

	user := api.Group("user")
	user.Post("/create", di.GetUserHandler().CreateUser)
	user.Put("/:id", di.GetUserHandler().UpdateUser)
	user.Delete("/:id", di.GetUserHandler().DeleteUser)
	user.Get("/", di.GetUserHandler().FindAll)
	user.Get("/:id", di.GetUserHandler().FindById)
	user.Post("/address", di.GetUserHandler().CreateAddress)
	user.Put("/address/:id", di.GetUserHandler().UpdateAddress)
	user.Delete("/address/:id", di.GetUserHandler().DeleteAddress)
	user.Post("/contact", di.GetUserHandler().CreateContact)
	user.Put("/contact/:id", di.GetUserHandler().UpdateContact)
	user.Delete("/contact/:id", di.GetUserHandler().DeleteContact)

	service := api.Group("service")
	service.Post("/create", di.GetServiceHandler().Create)
	service.Put("/:id", di.GetServiceHandler().Update)
	service.Delete("/:id", di.GetServiceHandler().Delete)
	service.Get("/", di.GetServiceHandler().FindAll)
	service.Get("/:id", di.GetServiceHandler().FindOne)

	supplier := api.Group("/supplier")
	supplier.Post("/create", di.GetSupplierHandler().Create)
	supplier.Put("/:id", di.GetSupplierHandler().Update)
	supplier.Delete("/:id", di.GetSupplierHandler().Delete)
	supplier.Get("/", di.GetSupplierHandler().FindAll)
	supplier.Get("/:id", di.GetSupplierHandler().FindOne)

	orderServiceStatus := api.Group("order-service-status")
	orderServiceStatus.Post("/create", di.GetOrderStatusHandler().Create)
	orderServiceStatus.Put("/:id", di.GetOrderStatusHandler().Update)
	orderServiceStatus.Delete("/:id", di.GetOrderStatusHandler().Delete)
	orderServiceStatus.Get("/", di.GetOrderStatusHandler().FindAll)
	orderServiceStatus.Get("/:id", di.GetOrderStatusHandler().FindOne)

	orderService := api.Group("order-service")
	orderService.Post("/create", di.GetOrderServiceHandler().Create)
	orderService.Put("/:id", di.GetOrderServiceHandler().Update)
	orderService.Get("/:id", di.GetOrderServiceHandler().FindOne)
	orderService.Get("/", di.GetOrderServiceHandler().FindAll)

	stock := api.Group("stock")
	stock.Post("/create", di.GetStockHandler().Add)
	stock.Get("/movements/:productId", di.GetStockHandler().FindStock)

}
