package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"safpass-api/services"
)

type NotificationHandler struct {
	SubscriptionService *services.SubscriptionService
}

func NewNotificationHandler(subscriptionService *services.SubscriptionService) *NotificationHandler {
	return &NotificationHandler{
		SubscriptionService: subscriptionService,
	}
}

func (h *NotificationHandler) PaymentNotification(c *fiber.Ctx) error {
	var notification map[string]interface{}
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	log.Println("Payment notification received:", notification)

	// Extract relevant data from notification
	orderID := notification["order_id"].(string)
	transactionStatus := notification["transaction_status"].(string)

	// Update subscription status based on transaction status
	err := h.SubscriptionService.UpdateSubscriptionStatus(orderID, transactionStatus)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *NotificationHandler) RecurringNotification(c *fiber.Ctx) error {
	var notification map[string]interface{}
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	log.Println("Recurring notification received:", notification)

	// Handle the recurring notification here

	return c.SendStatus(fiber.StatusOK)
}

func (h *NotificationHandler) PayAccountNotification(c *fiber.Ctx) error {
	var notification map[string]interface{}
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	log.Println("Pay account notification received:", notification)

	// Handle the pay account notification here

	return c.SendStatus(fiber.StatusOK)
}

func (h *NotificationHandler) FinishRedirect(c *fiber.Ctx) error {
	orderID := c.Query("order_id")
	statusCode := c.Query("status_code")
	transactionStatus := c.Query("transaction_status")

	log.Println("Finish redirect received:", orderID, statusCode, transactionStatus)

	// Handle the finish redirect here
	// You can redirect the user to a success page or handle the order status

	return c.Redirect("/payment-success") // Redirect to your success page
}

func (h *NotificationHandler) UnfinishedRedirect(c *fiber.Ctx) error {
	orderID := c.Query("order_id")
	statusCode := c.Query("status_code")
	transactionStatus := c.Query("transaction_status")

	log.Println("Unfinish redirect received:", orderID, statusCode, transactionStatus)

	// Handle the unfinish redirect here
	// You can redirect the user to an unfinish page or handle the order status

	return c.Redirect("/payment-unfinished") // Redirect to your unfinished page
}

func (h *NotificationHandler) ErrorRedirect(c *fiber.Ctx) error {
	orderID := c.Query("order_id")
	statusCode := c.Query("status_code")
	transactionStatus := c.Query("transaction_status")

	log.Println("Error redirect received:", orderID, statusCode, transactionStatus)

	// Handle the error redirect here
	// You can redirect the user to an error page or handle the order status

	return c.Redirect("/payment-error") // Redirect to your error page
}
