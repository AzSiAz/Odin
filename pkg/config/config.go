// Package config Should be used when you have config to be availble accross the application without having
// to get & check env variable
package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port      int
	DBAdress  string
	Host      string
	JWTSecret string
}

func New() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(fmt.Errorf("PORT could not be converted to int: %w", err))
	}

	dbAdress := os.Getenv("DB_ADRESS")
	if dbAdress == "" {
		panic(fmt.Errorf("DB_ADRESS is empty: %w", err))
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic(fmt.Errorf("JWT_SECRET is empty: %w", err))
	}

	return &Config{
		Port:      port,
		DBAdress:  dbAdress,
		JWTSecret: jwtSecret,
	}
}
