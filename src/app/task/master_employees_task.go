package task

import (
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

func (m *masterEmployeesTask) FindAll() (response []*model.MasterEmployeesModel, err error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)
	employeeService := employees.NewFindAllMasterEmployeesService(employeeRepo)

	response, err = employeeService.FindAll()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (m *masterEmployeesTask) Add(request *model.MasterEmployeesModel) (*model.MasterEmployeesModel, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)
	employeeService := employees.NewAddMasterEmployeesService(employeeRepo)

	err := employeeService.Add(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (m *masterEmployeesTask) Edit(request *model.MasterEmployeesModel) (*model.MasterEmployeesModel, error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)
	employeeService := employees.NewEditMasterEmployeesService(employeeRepo)

	err := employeeService.Edit(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}
