package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Env       string        `env:"ENV"`
	Mongo     MongoDBConfig `json:"mongo"`
	Redis     RedisConfig   `json:"redis"`
	Host      string        `env:"APP_HOST"`
	Port      string        `env:"APP_PORT"`
	JWTSECRET string        `env:"JWT_SECRET"`
}

func LoadConfig() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	environmentPath := filepath.Join(currentPath, ".env")

	if err := godotenv.Load(environmentPath); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

func GetConfig() Config {
	return Config{
		Env:       os.Getenv("ENV"),
		Mongo:     GetMongoDBConfig(),
		Redis:     GetRedisConfig(),
		Host:      os.Getenv("APP_HOST"),
		Port:      os.Getenv("APP_PORT"),
		JWTSECRET: os.Getenv("JWT_SECRET"),
	}
}
