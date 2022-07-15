package employees

import (
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/entity"
)

type FindAllMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewFindAllMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *FindAllMasterEmployeesService {
	return &FindAllMasterEmployeesService{
		Repository: repository,
	}
}

func (m *FindAllMasterEmployeesService) FindAll() (response []*model.MasterEmployeesResponseModel, err error) {

	result, err := m.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	response = m.cleanseData(result)

	return response, nil
}

func (m *FindAllMasterEmployeesService) cleanseData(masterData []*entity.MasterEmployeesEntity) (data []*model.MasterEmployeesResponseModel) {
	for _, row := range masterData {
		row.JoinDate = row.JoinDate[:10]
		row.DOB = row.DOB[:10]

		data = append(data, (*model.MasterEmployeesResponseModel)(row))
	}

	return data
}
