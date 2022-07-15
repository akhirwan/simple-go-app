package handler

import (
	"fmt"
	"simple-go-app/src/app/task"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/helper"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/config"
	"strconv"

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

	responses, err := task.FindAll()
	exception.PanicIfNeeded(err)

	helper.MessageOK = fmt.Sprintf("Get %v row data successed", len(responses))

	if len(responses) > 1 {
		helper.MessageOK = fmt.Sprintf("Get %v rows data successed", len(responses))
	}

	return helper.ResponseOK(c, responses)
}

func (m *MasterEmployeesHandler) Show(c *fiber.Ctx) error {
	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	httpStatus, response, err := task.Show(c.Params("id"))
	exception.PanicIfNeeded(err)

	if httpStatus == 404 {
		helper.MessageOK = fmt.Sprintf("ID %s is not found", c.Params("id"))
		return helper.ResponseNotFound(c, nil)
	}

	helper.MessageOK = fmt.Sprintf("ID %s is found", c.Params("id"))
	return helper.ResponseOK(c, response)
}

func (m *MasterEmployeesHandler) Add(c *fiber.Ctx) error {
	var request model.MasterEmployeesRequestModel

	err := c.BodyParser(&request)
	exception.PanicIfBadRequest(err)

	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	response, err := task.Add(&request)
	exception.PanicIfNeeded(err)

	helper.MessageOK = "1 Data recorded"
	return helper.ResponseOK(c, response)
}

func (m *MasterEmployeesHandler) Edit(c *fiber.Ctx) error {
	var request model.MasterEmployeesRequestModel

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

func (m *MasterEmployeesHandler) Activate(c *fiber.Ctx) error {
	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	httpStatus, isActive, err := task.Activate(c.Params("id"))
	exception.PanicIfNeeded(err)

	if httpStatus == 404 {
		helper.MessageOK = fmt.Sprintf("ID %s is not found", c.Params("id"))
		return helper.ResponseNotFound(c, nil)
	}

	helper.MessageOK = fmt.Sprintf("Data with ID %v is updated", c.Params("id"))
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	return helper.ResponseOK(c, &model.MasterEmployeesActivateResponseModel{
		ID:       id,
		IsActive: isActive,
	})
}

func (m *MasterEmployeesHandler) Delete(c *fiber.Ctx) error {
	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	httpStatus, isDeleted, err := task.Delete(c.Params("id"))
	exception.PanicIfNeeded(err)

	if httpStatus == 404 {
		helper.MessageOK = fmt.Sprintf("ID %s is not found", c.Params("id"))
		return helper.ResponseNotFound(c, nil)
	}

	helper.MessageOK = fmt.Sprintf("Data with ID %v is deleted", c.Params("id"))
	if !isDeleted {
		helper.MessageOK = fmt.Sprintf("Data with ID %v is reappeared", c.Params("id"))
	}

	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	return helper.ResponseOK(c, &model.MasterEmployeesDeleteResponseModel{
		ID:        id,
		IsDeleted: isDeleted,
	})
}

func (m *MasterEmployeesHandler) Remove(c *fiber.Ctx) error {
	task := task.NewMasterEmployeesTask(&task.Task{}, m.Config)
	httpStatus, err := task.Remove(c.Params("id"))
	exception.PanicIfNeeded(err)

	if httpStatus == 404 {
		helper.MessageOK = fmt.Sprintf("ID %s is not found", c.Params("id"))
		return helper.ResponseNotFound(c, nil)
	}

	helper.MessageOK = fmt.Sprintf("ID %s cannot be found in the future", c.Params("id"))
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)
	return helper.ResponseOK(c, &model.MasterEmployeesRemoveResponseModel{
		ID:        id,
		IsRemoved: true,
	})
}
