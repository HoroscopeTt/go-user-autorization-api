package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/horoscope/go-hexagonal-template/adapters/util"
	"github.com/horoscope/go-hexagonal-template/config"
)

func main() {
	cfg := config.LoadConfig()

	mongoClient, err := util.InitMongoDBClient(cfg.MongoDBURI, cfg.MongoDBTimeout)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		} else {
			log.Println("Disconnected from MongoDB.")
		}
	}()

	app := fiber.New()
	app.Use(logger.New())

	log.Printf("Bookstore API Service started on :%s", cfg.ServerPort)
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}