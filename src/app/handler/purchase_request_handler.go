package handler

import (
	"simple-procurement/src/app/task"
	"simple-procurement/src/domain/helper"
	"simple-procurement/src/infrastructure/config"

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
