// Package handler collection of http handler for router
package handler

import (
	"odin/pkg/config"
	"odin/pkg/database"
	"odin/pkg/user"
)

type Handler struct {
	db          *database.StormDB
	config      *config.Config
	userService *user.UserService
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
}

func New(db *database.StormDB, config *config.Config) *Handler {
	userService := user.NewUserService(db)

	return &Handler{
		db:          db,
		config:      config,
		userService: userService,
	}
}
