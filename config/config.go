package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPass             string
	DBName             string
	JWTSecret          string
	JWTExpirationHours int
}

func LoadConfig() *Config {
	var config Config
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPass = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.JWTSecret = os.Getenv("JWT_SECRET")
	config.JWTExpirationHours, _ = strconv.Atoi(os.Getenv("JWT_EXP_HRS"))
	return &config
}
