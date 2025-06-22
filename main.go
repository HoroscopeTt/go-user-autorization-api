package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/horoscope/go-hexagonal-template/config"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()
	app.Use(logger.New())

	log.Printf("Bookstore API Service started on :%s", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))

}