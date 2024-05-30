package repository

import (
	"context"
	"github.com/google/uuid"
	"safpass-api/database"
	"safpass-api/models"
)

func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := database.DB.QueryRow(context.Background(), `
		SELECT u.id, u.first_name, u.last_name, u.username, u.email, u.password, u.role_id, r.name, u.created_at, u.updated_at
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1`, id).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Role.ID, &user.Role.Name, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
