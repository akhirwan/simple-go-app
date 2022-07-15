package task

import (
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/helper"
	"simple-go-app/src/domain/master/employees"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/config"
	"simple-go-app/src/infrastructure/db"
	"simple-go-app/src/infrastructure/repository"

	"github.com/jmoiron/sqlx"
)

type masterEmployeesTask struct {
	Task *Task
	DB   *sqlx.DB
}

func NewMasterEmployeesTask(
	task *Task,
	config config.Configuration,
) *masterEmployeesTask {
	return &masterEmployeesTask{
		Task: task,
		DB:   db.NewMySQLDBConnection(helper.CreateMySQLConfig(config), config),
	}
}

func (m *masterEmployeesTask) FindAll() (responses []*model.MasterEmployeesResponseModel, err error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)

	employeeService := employees.NewFindAllMasterEmployeesService(employeeRepo)
	responses, err = employeeService.FindAll()
	exception.PanicIfNeeded(err)

	return responses, nil
}

func (m *masterEmployeesTask) Show(id string) (int, *model.MasterEmployeesResponseModel, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)

	employeeService := employees.NewShowMasterEmployeesService(employeeRepo)
	httpStatus, response, err := employeeService.Show(id)
	exception.PanicIfNeeded(err)

	return httpStatus, response, nil
}

func (m *masterEmployeesTask) Add(request *model.MasterEmployeesRequestModel) (int, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)

	employeeService := employees.NewAddMasterEmployeesService(employeeRepo)
	response, err := employeeService.Add(request)
	exception.PanicIfNeeded(err)

	return response, nil
}

func (m *masterEmployeesTask) Edit(request *model.MasterEmployeesRequestModel) (int, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)

	employeeService := employees.NewEditMasterEmployeesService(employeeRepo)
	httpStatus, err := employeeService.Edit(request)
	exception.PanicIfNeeded(err)

	return httpStatus, nil
}

func (m *masterEmployeesTask) Activate(id string) (int, bool, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)

	employeeService := employees.NewActivateMasterEmployeesService(employeeRepo)
	httpStatus, isActive, err := employeeService.Activate(id)
	exception.PanicIfNeeded(err)

	return httpStatus, isActive, nil
}

func (m *masterEmployeesTask) Delete(id string) (int, bool, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)

	employeeService := employees.NewDeleteMasterEmployeesService(employeeRepo)
	httpStatus, isActive, err := employeeService.Delete(id)
	exception.PanicIfNeeded(err)

	return httpStatus, isActive, nil
}
