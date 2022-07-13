package task

import (
	"log"
	"simple-procurement/src/domain/helper"
	"simple-procurement/src/infrastructure/config"
	"simple-procurement/src/infrastructure/db"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type purchaseRequestTask struct {
	Task *Task
	DB   *sqlx.DB
}

func NewPurchaseRequestTask(Task *Task, config config.Configuration) *purchaseRequestTask {
	return &purchaseRequestTask{
		Task: Task,
		DB:   db.NewMySQLDBConnection(helper.CreateMySQLConfig(config), config),
	}
}

func (p *purchaseRequestTask) Execute(c *fiber.Ctx) error {
	log.Println("bakekoxxxx")
	return c.SendString("bakekoxxxx")
}
