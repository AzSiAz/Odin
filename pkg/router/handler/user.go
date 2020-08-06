package handler

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetMyself(c *fiber.Ctx) {
	userInfo := h.ExtractUserInfoFromToken(c)

	user, err := h.userService.GetUserByID(userInfo.ID)
	if err != nil {
		logrus.WithField("jwt", userInfo).Errorln(fmt.Errorf("error trying to get user by id: %w", err))
		c.Status(fiber.StatusInternalServerError).JSON(Response{
			Message: "Could not get user information, this should not happen",
		})
	}

	c.Status(fiber.StatusOK).JSON(user)
}
