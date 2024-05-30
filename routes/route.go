package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"safpass-api/handlers"
	"safpass-api/services"
)

func SetupRoutes(app *fiber.App) {
	// Default config
	app.Static("/public", "./public")

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:  "Origin, Content-Type, Accept",
		ExposeHeaders: "Content-Length",
	}))

	// Initialize services
	authService := services.NewAuthService()
	midtransService := services.NewMidtransService()
	subscriptionService := services.NewSubscriptionService(midtransService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)
	notificationHandler := handlers.NewNotificationHandler(subscriptionService)

	// Create a new route group with the prefix "/api/v1"
	api := app.Group("/api/v1")

	authRoutes := api.Group("/auth")
	{
		authRoutes.Post("/register", authHandler.CreateUser)
		authRoutes.Post("/login", authHandler.AuthenticateUser)
	}

	subscriptionRoutes := api.Group("/subscription")
	{
		subscriptionRoutes.Post("/new", subscriptionHandler.SubscribeUser)
		subscriptionRoutes.Get("/user", subscriptionHandler.GetUserSubscription)
	}

	notificationRoutes := api.Group("/notification")
	{
		notificationRoutes.Post("/callback", notificationHandler.PaymentNotification)
	}
}
