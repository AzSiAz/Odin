package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func AuthorizationHeaderCheck(key string) func(*fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		if auth := ctx.Get("Authorization"); !(key == auth) {
			logrus.WithField("token", auth).Info("Endpoint called with wrong authorization token")
			ctx.SendStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
