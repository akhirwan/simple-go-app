package handler

import (
	"simple-go-app/src/app/task"
	"simple-go-app/src/domain/helper"
	"simple-go-app/src/infrastructure/config"

	"github.com/gofiber/fiber/v2"
)

type PurchaseRequestHandler struct {
	Config config.Configuration
}

func NewPurchaseRequestHandler(config config.Configuration) PurchaseRequestHandler {
	return PurchaseRequestHandler{
		Config: config,
	}
}

func (p *PurchaseRequestHandler) GetAll(c *fiber.Ctx) error {

	task := task.NewPurchaseRequestTask(&task.Task{}, p.Config)
	response := task.Execute(c)

	return helper.ResponseOK(c, response.Error())
}
