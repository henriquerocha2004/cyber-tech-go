package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/di"
	"github.com/henriquerocha2004/cyber-tech-go/internal/infra/routes"
	"github.com/spf13/viper"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	viper.SetConfigName("env")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error in read file configuration %w", err))
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	routes.Register(app)
	port := viper.GetString("webserver.port")
	//go launchListeners()
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}

func launchListeners() {
	//instance listeners
	di := &di.DependencyContainer{}
	orderServiceListener := di.GetOrderServiceListenEvents()
	orderServiceListener.GetEvents()
}
