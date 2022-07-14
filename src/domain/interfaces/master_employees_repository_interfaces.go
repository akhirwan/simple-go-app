package interfaces

import (
	"simple-go-app/src/domain/model"
	"time"
)

type MasterEmployeesRepositoryInterfaces interface {
	Insert(data []*model.MasterEmployeesModel) error
	FindAll() (data []*model.MasterEmployeesModel, err error)
	FindLastID(date time.Time) (result int64, err error)
	IfExistByID(id string) (bool, error)
}
