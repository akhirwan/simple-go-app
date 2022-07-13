package app

import (
	"simple-procurement/src/app/webapi"
	"simple-procurement/src/infrastructure/config"
	"simple-procurement/src/infrastructure/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	ID        string
	Name      string
	Version   string
	Config    config.Configuration
	WebServer web.Server
}

func NewApp() *App {
	configuration := config.New()
	return &App{
		ID:      configuration.Get("APP_ID"),
		Name:    configuration.Get("APP_NAME"),
		Version: configuration.Get("APP_VERSION"),
		Config:  configuration,
	}
}

func (apps *App) Run() *App {
	var app *fiber.App

	app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	api := webapi.NewAPIRouter(app, apps.Config)
	apps.WebServer = web.NewHTTPServer(apps.Config.Get("APP_PORT"), api.Route())

	apps.WebServer.Listen()

	return apps
}
