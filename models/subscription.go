package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	ID        int       `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	PlanID    int       `json:"plan_id" db:"plan_id"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Status    string    `json:"status" db:"status"`
	OrderID   string    `json:"order_id" db:"order_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
