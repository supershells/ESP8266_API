package configs

import (
	"os"
)

type MongoDBConfig struct {
	URI         string `env:"MONGODB_URI"`
	MongoDBName string `env:"MONGODB_NAME"`
}

func GetMongoDBConfig() MongoDBConfig {
	return MongoDBConfig{
		URI:         os.Getenv("MONGODB_URI"),
		MongoDBName: os.Getenv("MONGODB_NAME"),
	}
}
