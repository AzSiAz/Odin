package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Jwt string `json:"jwt"`
}

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) Login(c *fiber.Ctx) {
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: "Error on login request",
		})
		return
	}

	user, err := h.userService.GetUserByEmail(input.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(Response{
			Message: "Internal Server Error",
		})
		return
	}

	log.WithField("user", user).Info("Found User, checking password")
	if !h.userService.CheckPasswordHash(input.Password, user.Password) {
		c.Status(fiber.StatusUnauthorized).JSON(Response{
			Message: "Could not log you in",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	t, err := token.SignedString([]byte(h.config.JWTSecret))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(Response{
			Message: "Could not sign token, try again",
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(Response{
		Message: "Success login",
		Data:    LoginResponse{Jwt: t},
	})
}

func (h *Handler) SignUp(c *fiber.Ctx) {
	var input SignUpInput
	if err := c.BodyParser(&input); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: "Error on signup request",
		})
		return
	}

	user, err := h.userService.CreateUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(Response{
			Message: "Error creating user",
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(Response{
		Message: "User created",
		Data:    user,
	})
}
