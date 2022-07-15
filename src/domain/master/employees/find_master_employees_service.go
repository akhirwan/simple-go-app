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

func (m *FindAllMasterEmployeesService) FindAll() (responses []*model.MasterEmployeesResponseModel, err error) {

	result, err := m.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses = m.cleanseData(result)

	return responses, nil
}

func (m *FindAllMasterEmployeesService) cleanseData(masterData []*entity.MasterEmployeesEntity) (data []*model.MasterEmployeesResponseModel) {
	// pp.Println(len(masterData))
	for _, row := range masterData {
		// for _, rowx := range data {
		// row.JoinDate = row.JoinDate[:10]
		// row.DOB = row.DOB[:10]

		data = append(data, &model.MasterEmployeesResponseModel{
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
		})
		// }
	}

	return data
}
