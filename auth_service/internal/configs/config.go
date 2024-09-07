package configs

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DB_DSN      string
	GRPCAddress string
	HTTPAddress string
	JWTSecret   string
	JWTExpiry   int
}

func LoadConfig() *Config {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Panicln("Error loading .env")
	// }

	JWTExpiry, err := strconv.Atoi(os.Getenv("JWTEXPIRY"))
	if err != nil {
		log.Fatalf("Error ocured in config: %v", err)
	}

	var config Config = Config{
		DB_DSN:      os.Getenv("dsn"),
		GRPCAddress: ":" + os.Getenv("GRPCAddress"),
		HTTPAddress: ":" + os.Getenv("HTTPAddress"),
		JWTSecret:   os.Getenv("JWTSECRET"),
		JWTExpiry:   JWTExpiry,
	}

	var configPtr *Config = &config

	return configPtr

}
