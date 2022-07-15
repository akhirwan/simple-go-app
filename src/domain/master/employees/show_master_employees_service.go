package employees

import (
	"net/http"
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/entity"
)

type ShowMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewShowMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *ShowMasterEmployeesService {
	return &ShowMasterEmployeesService{
		Repository: repository,
	}
}

func (m *ShowMasterEmployeesService) Show(id string) (int, *model.MasterEmployeesResponseModel, error) {
	result, err := m.Repository.FindByID(id)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if result == nil {
		return http.StatusNotFound, nil, nil
	}

	return http.StatusOK, m.cleanseData(result), nil
}

func (m *ShowMasterEmployeesService) cleanseData(row *entity.MasterEmployeesEntity) *model.MasterEmployeesResponseModel {
	return &model.MasterEmployeesResponseModel{
		ID:       row.ID,
		Name:     row.Name,
		DeptID:   row.DeptID,
		Level:    row.Level,
		JoinDate: row.JoinDate[:10],
		IsActive: row.IsActive,
		Address:  row.Address,
		Email:    row.Email,
		Phone:    row.Phone,
		DOB:      row.DOB[:10],
		POB:      row.POB,
	}
}
