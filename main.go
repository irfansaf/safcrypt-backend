package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"safpass-api/configs"
	"safpass-api/database"
	"safpass-api/routes"
	"safpass-api/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.LoadConfig()

	database.Migrate()
	database.Init(config)
	utils.InitRedis()

	app := fiber.New()
	routes.SetupRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
