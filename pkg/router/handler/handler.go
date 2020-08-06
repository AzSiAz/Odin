// Package handler collection of http handler for router
package handler

import (
	"odin/pkg/config"
	"odin/pkg/database"
	"odin/pkg/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
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

type UserFromJWT struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func New(db *database.StormDB, config *config.Config) *Handler {
	userService := user.NewUserService(db)

	return &Handler{
		db:          db,
		config:      config,
		userService: userService,
	}
}

func (h *Handler) ExtractUserInfoFromToken(c *fiber.Ctx) *UserFromJWT {
	token := c.Locals("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	id := uint64(claims["id"].(float64))
	username := claims["username"].(string)

	return &UserFromJWT{
		ID:       id,
		Username: username,
	}
}
