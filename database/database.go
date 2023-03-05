package database

import (
	"context"
	cfg "esp8266_api/configs"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MgDB MongoInstance
	RdDB *redis.Client
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

// ConnectMongo Returns the Mongo DB Instance
func ConnectMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.GetConfig().Mongo.URI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(cfg.GetConfig().Mongo.MongoDBName)

	if err != nil {
		log.Println(strings.Repeat("#", 40))
		log.Println("Could Not Establish Mongo DB Connection")
		log.Println(strings.Repeat("#", 40))

		log.Fatal(err)
	}

	log.Println(strings.Repeat("#", 40))
	log.Println("Connected To Mongo DB")
	log.Println(strings.Repeat("#", 40))

	MgDB = MongoInstance{
		Client: client,
		Db:     db,
	}
}

func ConnectRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.GetConfig().Redis.HOST, cfg.GetConfig().Redis.PORT),
		Password: "p@ssw0rd",
		DB:       0,
	})

	pong, err := client.Ping(client.Context()).Result()
	if len(pong) > 0 {
		fmt.Println("pong die")
	}

	if err != nil {
		log.Println(strings.Repeat("#", 40))
		log.Println("Could Not Establish Redis Connection")
		log.Println(strings.Repeat("#", 40))
		log.Fatal(err)
	}

	log.Println(strings.Repeat("#", 40))
	log.Printf("Connected To Redis: %s\n", pong)
	log.Println(strings.Repeat("#", 40))

	RdDB = client
}
