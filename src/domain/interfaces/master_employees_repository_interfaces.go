package interfaces

import (
	"simple-go-app/src/infrastructure/entity"
	"time"
)

type MasterEmployeesRepositoryInterfaces interface {
	Record(data []*entity.MasterEmployeesEntity) error
	FindAll() (data []*entity.MasterEmployeesEntity, err error)
	FindByID(id string) (*entity.MasterEmployeesEntity, error)
	FindExistingByID(id string) (*entity.MasterEmployeesEntity, error)
	FindLastID(date time.Time) (result int64, err error)
	Delete(id string) error
}
