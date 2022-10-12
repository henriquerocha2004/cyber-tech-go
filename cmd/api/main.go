package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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
	routes.Register(app)
	port := viper.GetString("webserver.port")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}

func launchListeners() {

	for {
		fmt.Println("test")
	}

}
