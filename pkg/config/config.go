// Package config Should be used when you have config to be availble accross the application without having
// to get & check env variable
package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port      string
	DBAdress  string
	Host      string
	JWTSecret string
}

func New() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		panic(fmt.Errorf("PORT could not be converted to int: %s", port))
	}

	dbAdress := os.Getenv("DB_ADRESS")
	if dbAdress == "" {
		panic(fmt.Errorf("DB_ADRESS is empty: %s", dbAdress))
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic(fmt.Errorf("JWT_SECRET is empty: %s", jwtSecret))
	}

	return &Config{
		Port:      port,
		DBAdress:  dbAdress,
		JWTSecret: jwtSecret,
	}
}
