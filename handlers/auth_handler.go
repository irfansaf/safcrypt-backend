package handlers

import (
	"github.com/gofiber/fiber/v2"
	"safpass-api/configs"
	"safpass-api/models"
	"safpass-api/services"
	"safpass-api/utils"
	"strings"
)

type Handler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *Handler {
	return &Handler{
		AuthService: authService,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := h.AuthService.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
	})
}

func (h *Handler) AuthenticateUser(c *fiber.Ctx) error {
	var loginRequest models.User

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	loginRequest.Username = strings.ToLower(loginRequest.Username)

	user, err := h.AuthService.AuthenticateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	JWTSecretKey := configs.LoadConfig().JWTSecretKey

	token, err := utils.GenerateJWTToken(user.ID, JWTSecretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.StoreTokenInRedis(user.ID, token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User authenticated successfully",
		"data": fiber.Map{
			"access_token": token,
		},
	})
}
