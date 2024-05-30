package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"safpass-api/database"
	"safpass-api/models"
)

func CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (id, first_name, last_name, username, email, password)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.DB.Exec(
		context.Background(),
		query,
		user.ID, user.FirstName, user.LastName, user.Username, user.Email, user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(username string) (*models.User, error) {
	var user models.User
	var userID string
	var query string
	var err error

	query = `
		SELECT id, username, password, first_name, last_name, email
		FROM users
		WHERE username = $1`

	err = database.DB.QueryRow(
		context.Background(),
		query,
		username,
	).Scan(
		&userID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	user.ID, err = uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
