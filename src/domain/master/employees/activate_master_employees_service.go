package employees

import (
	"net/http"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/infrastructure/entity"
)

type ActivateMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewActivateMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *ActivateMasterEmployeesService {
	return &ActivateMasterEmployeesService{
		Repository: repository,
	}
}

func (m *ActivateMasterEmployeesService) Activate(id string) (httpStatus int, isActive bool, err error) {
	var data []*entity.MasterEmployeesEntity
	var row *entity.MasterEmployeesEntity

	row, err = m.Repository.FindByID(id)
	exception.PanicIfNeeded(err)

	if row == nil {
		return http.StatusNotFound, false, nil
	}

	if !row.IsActive {
		row.IsActive = true
	} else {
		row.IsActive = false
	}

	data = append(data, row)

	err = m.Repository.Record(data)
	exception.PanicIfNeeded(err)

	return http.StatusOK, row.IsActive, nil
}
