package handler

import (
	"fmt"
	"simple-go-app/src/app/task"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/helper"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/config"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MasterEmployeesHandler struct {
	Config config.Configuration
}

func NewMasterEmployeesHandler(config config.Configuration) MasterEmployeesHandler {
	return MasterEmployeesHandler{
		Config: config,
	}
}

func (m *MasterEmployeesHandler) FindAll(c *fiber.Ctx) error {
	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)

	response, err := task.FindAll()
	if err != nil {
		return err
	}

	helper.MessageOK = fmt.Sprintf("Get %v row data successed", len(response))

	if len(response) > 1 {
		helper.MessageOK = fmt.Sprintf("Get %v rows data successed", len(response))
	}

	return helper.ResponseOK(c, response)
}

func (m *MasterEmployeesHandler) Add(c *fiber.Ctx) error {
	var request model.MasterEmployeesModel

	err := c.BodyParser(&request)
	exception.PanicIfBadRequest(err)

	joinDate, _ := time.Parse("2006-01-02", request.JoinDate)
	request.ID, _ = strconv.ParseInt(joinDate.Format("20060102")+"000", 10, 64)

	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	response, err := task.Add(&request)
	exception.PanicIfNeeded(err)

	helper.MessageOK = "1 Data recorded"
	return helper.ResponseOK(c, response)
}

func (m *MasterEmployeesHandler) Edit(c *fiber.Ctx) error {
	var request model.MasterEmployeesModel

	err := c.BodyParser(&request)
	exception.PanicIfBadRequest(err)

	request.ID, _ = strconv.ParseInt(c.Params("id"), 10, 64)

	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	status, err := task.Edit(&request)
	exception.PanicIfNeeded(err)

	if status == 404 {
		helper.MessageOK = fmt.Sprintf("ID %v is not found", request.ID)
		return helper.ResponseNotFound(c, nil)
	}

	helper.MessageOK = fmt.Sprintf("Data with ID %v is updated", request.ID)
	return helper.ResponseOK(c, request)
}
