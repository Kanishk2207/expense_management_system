package repository

import (
	"database/sql"
	"user-service/internal/models"
)

type UserRepositoy struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoy {
	userRepo := UserRepositoy{DB: db}
	userRepoPtr := &userRepo
	return userRepoPtr
}

func (r *UserRepositoy) CreateUser(user *models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (user_id, username, first_name, last_name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.UserID, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

func (r *UserRepositoy) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := r.DB.QueryRow(
		"SELECT user_id, username, first_name, last_name, email, password FROM users WHERE email = ?", email,
	).Scan(
		&user.UserID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}
