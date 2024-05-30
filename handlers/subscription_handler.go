package handlers

import (
	"github.com/gofiber/fiber/v2"
	"safpass-api/configs"
	"safpass-api/services"
	"safpass-api/utils"
)

type SubscriptionHandler struct {
	SubscriptionService *services.SubscriptionService
}

func NewSubscriptionHandler(subscriptionService *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		SubscriptionService: subscriptionService,
	}
}

func (h *SubscriptionHandler) SubscribeUser(c *fiber.Ctx) error {
	token, err := utils.ExtractTokenFromHeader(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID, err := utils.GetUserIDFromToken(token, configs.LoadConfig().JWTSecretKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	type Request struct {
		PlanID int `json:"plan_id"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	subscription, transactionResp, err := h.SubscriptionService.SubscribeUser(userID, req.PlanID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":      true,
		"message":      "Subscription initiated, please complete payment",
		"subscription": subscription,
		"payment_url":  transactionResp.RedirectURL,
	})
}

func (h *SubscriptionHandler) GetUserSubscription(c *fiber.Ctx) error {
	token, err := utils.ExtractTokenFromHeader(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID, err := utils.GetUserIDFromToken(token, configs.LoadConfig().JWTSecretKey)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subscription, err := h.SubscriptionService.GetUserSubscription(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if subscription == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No active subscription found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": subscription,
	})
}
