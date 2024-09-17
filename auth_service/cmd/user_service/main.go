package main

import (
	"fmt"
	"net/http"
	"user-service/internal/configs"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status": "up"}`)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	config := configs.LoadConfig()
	// dsn := "dev_user:Kanishk_22@tcp(mysql-service:3306)/auth_service_db"
	// db.InitDB(dsn)
	db.InitDB(config.DB_DSN)
	defer db.DB.Close()

	http.HandleFunc("/health", healthCheckHandler)
	port := config.HTTPAddress
	fmt.Printf("Starting server on port %s...\n", port)

	userRepo := repository.NewUserRepository(db.DB)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	http.HandleFunc("/signup", authHandler.Signup)
	http.HandleFunc("/login", authHandler.Login)

	http.ListenAndServe(port, nil)
}
