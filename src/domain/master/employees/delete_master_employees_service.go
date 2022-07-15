package employees

import (
	"net/http"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/infrastructure/entity"
)

type DeleteMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewDeleteMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *DeleteMasterEmployeesService {
	return &DeleteMasterEmployeesService{
		Repository: repository,
	}
}

func (m *DeleteMasterEmployeesService) Delete(id string) (httpStatus int, isDeleted bool, err error) {
	var data []*entity.MasterEmployeesEntity

	row, err := m.Repository.FindExistingByID(id)
	exception.PanicIfNeeded(err)

	if row == nil {
		return http.StatusNotFound, false, nil
	}

	if row.IsDeleted {
		row.IsDeleted = false
	} else {
		row.IsDeleted = true
	}

	data = append(data, row)

	err = m.Repository.Record(data)
	exception.PanicIfNeeded(err)

	return http.StatusOK, row.IsDeleted, nil
}
