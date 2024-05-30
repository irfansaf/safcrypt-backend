package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"log"
	"safpass-api/models"
	"safpass-api/repository"
	"time"
)

type SubscriptionService struct {
	MidtransService *MidtransService
}

func NewSubscriptionService(midtransService *MidtransService) *SubscriptionService {
	return &SubscriptionService{
		MidtransService: midtransService,
	}
}

func (s *SubscriptionService) SubscribeUser(userID uuid.UUID, planID int) (*models.Subscription, *snap.Response, error) {
	log.Println("Subscribing user to plan")
	plan, err := repository.GetPlanByID(planID)
	if err != nil {
		return nil, nil, err
	}

	log.Println("Plan found")

	if plan == nil {
		return nil, nil, errors.New("plan not found")
	}

	log.Println("Creating transaction")

	user, err := repository.GetUserByID(userID)
	if err != nil {
		return nil, nil, err
	}

	log.Println("User found")

	customerDetails := &midtrans.CustomerDetails{
		FName: user.FirstName,
		LName: user.LastName,
		Email: user.Email,
		Phone: user.Phone,
	}

	log.Println("Customer details created")

	orderID := uuid.New().String()
	transactionResp, err := s.MidtransService.CreateTransaction(orderID, int64(plan.Price), customerDetails)
	if err != nil {
		return nil, nil, err
	}

	log.Println("Transaction created:", transactionResp)

	if transactionResp.Token == "" {
		return nil, nil, errors.New("payment failed")
	}

	log.Println("Payment initiated, creating pending subscription")

	subscription := &models.Subscription{
		UserID:    userID,
		PlanID:    plan.ID,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, plan.Duration),
		Status:    "pending",
		OrderID:   orderID,
	}

	err = repository.CreateSubscription(subscription)
	if err != nil {
		return nil, nil, err
	}

	log.Println("Pending subscription created")

	return subscription, transactionResp, nil
}

func (s *SubscriptionService) GetUserSubscription(userID uuid.UUID) (*models.Subscription, error) {
	return repository.GetUserSubscription(userID)
}

func (s *SubscriptionService) UpdateSubscriptionStatus(orderID string, transactionStatus string) error {
	log.Println("Updating subscription status for order:", orderID, "with status:", transactionStatus)

	subscription, err := repository.GetSubscriptionByOrderID(orderID)
	if err != nil {
		return err
	}

	if subscription == nil {
		return errors.New("subscription not found")
	}

	switch transactionStatus {
	case "settlement":
		subscription.Status = "active"
	case "pending":
		subscription.Status = "pending"
	case "cancel", "deny", "expire":
		subscription.Status = "inactive"
	default:
		return errors.New("unhandled transaction status")
	}

	err = repository.UpdateSubscription(subscription)
	if err != nil {
		return err
	}

	log.Println("Subscription status updated successfully")
	return nil
}
