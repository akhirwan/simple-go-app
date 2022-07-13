package interfaces

import "simple-go-app/src/domain/model"

type MasterEmployeesRepositoryInterfaces interface {
	FindAll() ([]*model.MasterEmployeesResponseModel, error)
}
