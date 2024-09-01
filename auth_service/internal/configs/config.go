package configs

import (
	"os"
)

type Config struct {
	DB_DSN      string
	GRPCAddress string
	HTTPAddress string
}

func LoadConfig() *Config {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Panicln("Error loading .env")
	// }

	var config Config = Config{
		DB_DSN:      os.Getenv("dsn"),
		GRPCAddress: ":" + os.Getenv("GRPCAddress"),
		HTTPAddress: ":" + os.Getenv("HTTPAddress"),
	}

	var configPtr *Config = &config

	return configPtr

}
