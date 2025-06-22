package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)



type AppConfig struct {
	MongoDBURI string
	MongoDBNAME string
	MongoDBTimeout int // in seconds
	ServerPort string
}


func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
		log.Println("SERVER_PORT not set, using default:", serverPort)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		// mongoURI = "8080"
		log.Println("MONGO_URI not set, using default:", mongoURI)
	}

	mongoName := os.Getenv("MONGO_DB_NAME")
	if mongoName == "" {
		// mongoURI = "8080"
		log.Println("MONGO_DB_NAME not set, using default:", mongoName)
	}

	var mongoTimeout int
	mongoTimeoutStr := os.Getenv("MONGO_DB_TIMEOUT")
	if mongoTimeoutStr == "" {
		log.Println("MONGO_DB_TIMEOUT not set, using default:", mongoTimeoutStr)
	} else {
		mongoTimeout, err = strconv.Atoi(mongoTimeoutStr)
		if err != nil {
			log.Println("Error")
		}
	}


	return &AppConfig{
		ServerPort:  serverPort,
		MongoDBURI:  mongoURI,
		MongoDBNAME: mongoName,
		MongoDBTimeout: mongoTimeout,
	}
}