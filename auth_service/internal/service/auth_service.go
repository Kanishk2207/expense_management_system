package service

import (
	"log"
	goerrors "user-service/internal/go_errors"
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/internal/utils"
)

type AuthService struct {
	repo *repository.UserRepositoy
}

func NewAuthService(r *repository.UserRepositoy) *AuthService {
	authService := AuthService{repo: r}
	authServicePrt := &authService
	return authServicePrt
}

func (s *AuthService) SignUp(username, firsName, lastName, email, password string) error {
	userId := utils.GetUuid()
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("Error occured while hashing: %v", err)
	}
	currentTime := utils.GetCurrentUnixTime()
	user := models.User{
		UserID:    userId,
		Username:  username,
		FirstName: firsName,
		LastName:  lastName,
		Email:     email,
		Password:  passwordHash,
		CreatedAt: int(currentTime),
		UpdatedAt: int(currentTime),
	}

	userPtr := &user

	exists, err := s.repo.CheckUserExists(user.Email, user.Username)
	if err != nil {
		return err
	}
	if exists {
		return goerrors.ErrUserAlreadyExists
	}

	err = s.repo.CreateUser(userPtr)
	if err != nil {
		log.Fatalf("Error occured while creating user: %v", err)
		return err
	}
	return nil
}
