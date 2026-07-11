package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return &Config{
		Port:   port,
		DBHost: dbHost,
		DBPort: dbPort,
		DBUser: dbUser,
		DBPass: dbPass,
		DBName: dbName,
	}
}
