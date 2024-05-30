package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"safpass-api/database"
	"safpass-api/models"
)

func CreateSubscription(subscription *models.Subscription) error {
	query := `
    INSERT INTO subscriptions (user_id, plan_id, start_date, end_date, status, order_id)
    VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.DB.Exec(
		context.Background(),
		query,
		subscription.UserID, subscription.PlanID, subscription.StartDate, subscription.EndDate, subscription.Status, subscription.OrderID,
	)
	if err != nil {
		return err
	}

	return nil
}

func GetPlanByID(planID int) (*models.Plan, error) {
	var plan models.Plan
	query := `SELECT id, name, description, price, duration FROM plans WHERE id = $1`
	err := database.DB.QueryRow(context.Background(), query, planID).Scan(
		&plan.ID, &plan.Name, &plan.Description, &plan.Price, &plan.Duration,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &plan, nil
}

func GetUserSubscription(userID uuid.UUID) (*models.Subscription, error) {
	var subscription models.Subscription
	query := `SELECT id, user_id, plan_id, start_date, end_date, status, order_id FROM subscriptions WHERE user_id = $1`
	err := database.DB.QueryRow(context.Background(), query, userID).Scan(
		&subscription.ID, &subscription.UserID, &subscription.PlanID, &subscription.StartDate, &subscription.EndDate, &subscription.Status, &subscription.OrderID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &subscription, nil
}

func GetSubscriptionByOrderID(orderID string) (*models.Subscription, error) {
	var subscription models.Subscription
	query := `SELECT id, user_id, plan_id, start_date, end_date, status, order_id FROM subscriptions WHERE order_id = $1`
	err := database.DB.QueryRow(context.Background(), query, orderID).Scan(
		&subscription.ID, &subscription.UserID, &subscription.PlanID, &subscription.StartDate, &subscription.EndDate, &subscription.Status, &subscription.OrderID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &subscription, nil
}

func UpdateSubscription(subscription *models.Subscription) error {
	query := `
    UPDATE subscriptions
    SET status = $1
    WHERE order_id = $2`
	_, err := database.DB.Exec(
		context.Background(),
		query,
		subscription.Status, subscription.OrderID,
	)
	if err != nil {
		return err
	}

	return nil
}
