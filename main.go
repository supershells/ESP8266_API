package main

import (
	cfg "esp8266_api/configs"
	"esp8266_api/controllers"
	db "esp8266_api/database"
	"esp8266_api/middlewares"
	"esp8266_api/repositories"
	"esp8266_api/services"
	"fmt"

	//"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	//"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Init Config & DB")
	cfg.LoadConfig()
	config := cfg.GetConfig()
	db.ConnectMongo()
	db.ConnectRedis()

	fmt.Println("Implementing Plug and Adapter")
	userRepositoryMongoDB := repositories.NewUserRepositoryMongoDB(db.MgDB.Client, db.MgDB.Db)
	userService := services.NewUserService(userRepositoryMongoDB)
	userHandler := controllers.NewUserHandler(userService)

	temperatureRepositoryMongDB := repositories.NewtemperatureRepositoryMongoDB(db.MgDB.Client, db.MgDB.Db)
	//temperatureService := service.NewTemperatureService(temperatureRepositoryMongDB)  // apdater no redis cache
	temperatureService := services.NewTemperatureServiceRedis(temperatureRepositoryMongDB, db.RdDB)
	temperatureHandler := controllers.NewTemperatureHandler(temperatureService)

	fmt.Println("Init Fiber App ")
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	fmt.Println("Init Routes") //todolist seperate to route.go
	//user
	app.Post("/user/register", userHandler.Register)
	app.Post("/user/login", userHandler.Login)
	app.Get("/user/list", middlewares.JwtMiddleware(), userHandler.LoadUserList)

	//temperature
	app.Get("/temperatures", temperatureHandler.GetTemperatures)
	app.Get("/temperature/:temperatureID", temperatureHandler.GetTemperature)
	app.Post("/temperature", temperatureHandler.NewTemperature)
	app.Delete("/temperature/delete/:temperatureID", temperatureHandler.DeleteTemperature)

	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
