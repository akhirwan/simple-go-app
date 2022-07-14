package webapi

import (
	"simple-go-app/src/app/handler"
	"simple-go-app/src/infrastructure/config"

	"github.com/gofiber/fiber/v2"
)

type APIRouter struct {
	App    *fiber.App
	Config config.Configuration
}

func NewAPIRouter(app *fiber.App, config config.Configuration) *APIRouter {
	return &APIRouter{
		App:    app,
		Config: config,
	}
}

func (a *APIRouter) Route() *fiber.App {

	a.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, UnderWorld!")
	})

	prHandler := handler.NewPurchaseRequestHandler(a.Config)
	a.App.Get("/request", prHandler.GetAll)

	masterHandler := handler.NewMasterEmployeesHandler(a.Config)
	a.App.Get("/employees", masterHandler.FindAll)
	a.App.Post("/employee", masterHandler.Add)

	return a.App
}
