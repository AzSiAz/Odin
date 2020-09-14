//go:generate pkger

package main

import (
	"odin/pkg/config"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Starting Odin")
	config := config.New()
	// db, err := database.New(config.DBAdress)
	// if err != nil {
	// panic(err)
	// }

	app := fiber.New()

	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	// app.Use(helmet.New())
	// app.Use(cors.New())

	// router.InitRouter(app, db, config)

	app.Use(filesystem.New(filesystem.Config{
		Root:         pkger.Dir("/web/dist/"),
		Index:        "index.html",
		NotFoundFile: "/index.html",
	}))

	logrus.Fatal(app.Listen(config.Port))
}
