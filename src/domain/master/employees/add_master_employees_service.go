package employees

import (
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
	"time"
)

type AddMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewAddMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *AddMasterEmployeesService {
	return &AddMasterEmployeesService{
		Repository: repository,
	}
}

func (m *AddMasterEmployeesService) Add(request *model.MasterEmployeesModel) error {
	var data []*model.MasterEmployeesModel
	var err error

	joinDate, _ := time.Parse("2006-01-02", request.JoinDate)
	request.ID, err = m.generateID(joinDate)
	if err != nil {
		return err
	}

	data = append(data, request)

	err = m.Repository.Insert(data)
	if err != nil {
		return err
	}

	return nil
}

func (m *AddMasterEmployeesService) generateID(date time.Time) (int64, error) {
	id, err := m.Repository.FindLastID(date)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}
