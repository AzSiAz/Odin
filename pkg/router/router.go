// Package router Handle router for fiber app
package router

import (
	"odin/pkg/config"
	"odin/pkg/database"
	"odin/pkg/router/handler"

	"github.com/gofiber/fiber"
)

// InitRouter Init backend routing for api
func InitRouter(app *fiber.App, db *database.StormDB, config *config.Config) {
	h := handler.New(db, config)
	api := app.Group("api")
	// webhook := app.Group("webhooks")

	api.Get("/", func(c *fiber.Ctx) {
		c.SendString("Hello, world !!")
	})

	api.Post("/login", h.Login)
	api.Post("/signup", h.SignUp)
}
