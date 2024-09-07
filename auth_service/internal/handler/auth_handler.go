package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"user-service/internal/dto"
	goerrors "user-service/internal/go_errors"
	"user-service/internal/service"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *service.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	authHandler := AuthHandler{authService: s, validator: validator.New()}
	authHandlerPtr := &authHandler
	return authHandlerPtr
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Mehtod not allowed", http.StatusMethodNotAllowed)
	}

	var req dto.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Input Format", http.StatusBadRequest)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		http.Error(w, "Request error: bad request", http.StatusBadRequest)
		log.Printf("Validation error: " + err.Error())
		return
	}

	err = h.authService.SignUp(req.Username, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, goerrors.ErrUserAlreadyExists) {
			http.Error(w, "User already exists", http.StatusConflict)
		} else {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}
