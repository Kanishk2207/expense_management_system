package repository

import (
	"database/sql"
)

type UserRepositoy struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoy {
	userRepo := UserRepositoy{DB: db}
	userRepoPtr := &userRepo
	return userRepoPtr
}

// func (r *UserRepositoy) CreateUser(user *models.User) error {
// 	_, err := r.DB.Exec()
// }
