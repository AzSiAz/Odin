// Package router Handle router for fiber app
package router

import (
	"odin/pkg/config"
	"odin/pkg/database"

	"github.com/gofiber/fiber"
)

// InitRouter Init backend routing for api
func InitRouter(app *fiber.App, db database.Database, config *config.Config) {
	api := app.Group("api")
	// webhook := app.Group("webhooks")

	// userService := user.NewUserService(db)

	api.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World")
	})
}
