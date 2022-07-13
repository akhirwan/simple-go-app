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

func (m *masterEmployeesTask) FindAll() (response []*model.MasterEmployeesResponseModel, err error) {
	employeeRepo := repository.NewMasterEmployeesRepository(m.DB)
	employeeService := employees.NewMasterEmployeesService(employeeRepo)

	response, err = employeeService.FindAll()
	if err != nil {
		return nil, err
	}

	return response, nil
}
