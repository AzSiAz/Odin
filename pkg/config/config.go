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
	Port              int
	DBAdress          string
	Host              string
	Auth0WebhookToken string
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

	auth0WebhookToken := os.Getenv("WEBHOOK_HEADER_AUTH_TOKEN")
	if auth0WebhookToken == "" {
		panic(fmt.Errorf("WEBHOOK_HEADER_AUTH_TOKEN is empty: %w", err))
	}

	return &Config{
		Port:              port,
		DBAdress:          dbAdress,
		Auth0WebhookToken: auth0WebhookToken,
	}
}
