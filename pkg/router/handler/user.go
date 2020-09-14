package handler

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetMyself(c *fiber.Ctx) error {
	userInfo := h.ExtractUserInfoFromToken(c)

	user, err := h.userService.GetUserByID(userInfo.ID)
	if err != nil {
		logrus.WithField("jwt", userInfo).Errorln(fmt.Errorf("error trying to get user by id: %w", err))
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Message: "Could not get user information, this should not happen",
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Message: "user data",
		Data:    user,
	})
}
