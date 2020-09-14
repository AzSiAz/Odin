// Package middlewares Collection of middleware used by odin
package middlewares

import (
	"odin/pkg/config"

	"github.com/dgrijalva/jwt-go"
	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected(config *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler:  jwtError,
		SigningKey:    []byte(config.JWTSecret),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}

// TODO as a param for Proctected to move it to handler package
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"message": "Missing or malformed JWT"})
	} else {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"message": "Invalid or expired JWT"})
	}
}
