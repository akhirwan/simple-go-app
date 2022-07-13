package employees

import (
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
)

type MasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *MasterEmployeesService {
	return &MasterEmployeesService{
		Repository: repository,
	}
}

func (m *MasterEmployeesService) FindAll() ([]*model.MasterEmployeesResponseModel, error) {
	var result []*model.MasterEmployeesResponseModel

	result, err := m.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	result = m.cleanseData(result)

	return result, nil
}

func (m *MasterEmployeesService) cleanseData(masterData []*model.MasterEmployeesResponseModel) []*model.MasterEmployeesResponseModel {
	var data []*model.MasterEmployeesResponseModel

	for _, row := range masterData {
		row.JoinDate = row.JoinDate[:10]
		row.DOB = row.DOB[:10]

		data = append(data, row)
	}

	return data
}
