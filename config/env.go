package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MongoConfig struct {
	Host           string
	Port           string
	Username       string
	Password       string
	DBName         string
	CollectionName string
}

func EnvMongo() MongoConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file")
	}

	return MongoConfig{
		Host:           os.Getenv("MONGO_HOST"),
		Port:           os.Getenv("MONGO_PORT"),
		Username:       os.Getenv("MONGO_USERNAME"),
		Password:       os.Getenv("MONGO_PASSWORD"),
		DBName:         os.Getenv("MONGO_DB_NAME"),
		CollectionName: os.Getenv("MONGO_COLLECTION_NAME"),
	}
}
