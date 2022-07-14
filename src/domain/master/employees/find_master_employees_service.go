package employees

import (
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
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

func (m *FindAllMasterEmployeesService) FindAll() (result []*model.MasterEmployeesModel, err error) {

	result, err = m.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	result = m.cleanseData(result)

	return result, nil
}

func (m *FindAllMasterEmployeesService) cleanseData(masterData []*model.MasterEmployeesModel) []*model.MasterEmployeesModel {
	var data []*model.MasterEmployeesModel

	for _, row := range masterData {
		row.JoinDate = row.JoinDate[:10]
		row.DOB = row.DOB[:10]

		data = append(data, row)
	}

	return data
}
