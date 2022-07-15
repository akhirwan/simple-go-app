package employees

import (
	"fmt"
	"net/http"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/entity"
	"time"
)

type EditMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewEditMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *EditMasterEmployeesService {
	return &EditMasterEmployeesService{
		Repository: repository,
	}
}

func (m *EditMasterEmployeesService) Edit(request *model.MasterEmployeesRequestModel) (httpStatus int, err error) {
	var data []*entity.MasterEmployeesEntity

	// isExist, err = m.Repository.IsExistByID(fmt.Sprintf("%v", request.ID))
	result, err := m.Repository.FindByID(fmt.Sprintf("%v", request.ID))
	exception.PanicIfNeeded(err)

	if result == nil {
		return http.StatusNotFound, nil
	}

	row := &entity.MasterEmployeesEntity{
		ID:         request.ID,
		Name:       request.Name,
		DeptID:     request.DeptID,
		Level:      request.Level,
		JoinDate:   request.JoinDate,
		IsActive:   result.IsActive,
		Address:    request.Address,
		Email:      request.Email,
		Phone:      request.Phone,
		DOB:        request.DOB,
		POB:        request.POB,
		IsDeleted:  false,
		CreatedAt:  result.CreatedAt,
		ModifiedAt: time.Now(),
	}

	data = append(data, row)

	err = m.Repository.Record(data)
	exception.PanicIfNeeded(err)

	return http.StatusOK, nil
}
