package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/configs"
	"user-service/internal/db"
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
	fmt.Println(config.DB_DSN)
	db.InitDB(config.DB_DSN)

	http.HandleFunc("/health", healthCheckHandler)
	port := config.HTTPAddress
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
