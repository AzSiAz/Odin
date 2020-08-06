//go:generate pkger

package main

import (
	"odin/pkg/config"
	"odin/pkg/database"
	"odin/pkg/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Starting Odin")
	config := config.New()
	db, err := database.New(config.DBAdress)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(middleware.Recover())
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())
	app.Use(middleware.Compress())
	app.Use(helmet.New())
	app.Use(cors.New())

	router.InitRouter(app, db, config)

	app.Use("/", middleware.FileSystem(middleware.FileSystemConfig{
		Root:   pkger.Dir("/web/dist"),
		Index:  "index.html",
		Browse: false,
	}))

	logrus.Fatal(app.Listen(config.Port))
}
