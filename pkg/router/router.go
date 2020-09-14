// Package router Handle router for fiber app
package router

import (
	"odin/pkg/config"
	"odin/pkg/database"
	"odin/pkg/router/handler"
	"odin/pkg/router/middlewares"

	fiber "github.com/gofiber/fiber/v2"
)

// InitRouter Init backend routing for api
func InitRouter(app *fiber.App, db *database.StormDB, config *config.Config) {
	h := handler.New(db, config)
	api := app.Group("api")
	// webhook := app.Group("webhooks")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world from Odin !!")
	})

	api.Post("/login", h.Login)
	api.Post("/signup", h.SignUp)

	protected := middlewares.Protected(config)

	api.Get("/user/me", protected, h.GetMyself)
}
