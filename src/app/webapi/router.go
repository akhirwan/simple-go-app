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

	employeesHandler := handler.NewMasterEmployeesHandler(a.Config)
	a.App.Get("/employees", employeesHandler.FindAll)
	a.App.Post("/employee", employeesHandler.Add)
	a.App.Put("/employee/:id", employeesHandler.Edit)
	a.App.Patch("employee/activate/:id", employeesHandler.Activate)

	// departmentsHandler

	// usersHandler

	// itemsHandler

	// categoriesHandler

	return a.App
}
