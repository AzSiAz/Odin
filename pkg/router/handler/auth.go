package handler

import (
	"errors"
	"fmt"
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
		log.WithField("err", err).Error("Error trying to parse user input")
		c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: "Error on login request",
		})
		return
	}

	if input.Email == "" || input.Password == "" {
		err := errors.New("one of the required field is empty or not present: email, password")
		c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: err.Error(),
		})
		return
	}

	user, err := h.userService.GetUserByEmail(input.Email)
	if err != nil {
		log.WithField("err", err).Error("Error trying to get user by email")
		c.Status(fiber.StatusInternalServerError).JSON(Response{
			Message: "Internal Server Error",
		})
		return
	}

	log.WithField("userId", user.ID).Info("Found User, checking password")
	if !h.userService.CheckPasswordHash(input.Password, user.Password) {
		log.Error("Error trying to check password")
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
		log.WithField("jwt", token).Error(fmt.Errorf("error trying to sign jwt token: %w", err))
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
		log.Error(fmt.Errorf("error trying to parse user input: %w", err))
		c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: "Error on signup request",
		})
		return
	}

	if input.Username == "" || input.Email == "" || input.Password == "" {
		err := errors.New("one of the required field is empty or not present: username, email, password")
		c.Status(fiber.StatusBadRequest).JSON(Response{
			Message: err.Error(),
		})
		return
	}

	user, err := h.userService.CreateUser(input.Username, input.Email, input.Password)
	if err != nil {
		log.WithFields(logrus.Fields{
			"email":    input.Email,
			"username": input.Username,
		}).Error(fmt.Errorf("error trying to create user: %w", err))
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
