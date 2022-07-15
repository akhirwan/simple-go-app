package employees

import (
	"net/http"
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
	"simple-go-app/src/infrastructure/entity"
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

func (m *AddMasterEmployeesService) Add(request *model.MasterEmployeesRequestModel) (response int, err error) {
	var data []*entity.MasterEmployeesEntity

	joinDate, _ := time.Parse("2006-01-02", request.JoinDate)
	request.ID, err = m.generateID(joinDate)
	if err != nil {
		return http.StatusBadRequest, err
	}

	row := &entity.MasterEmployeesEntity{
		ID:         request.ID,
		Name:       request.Name,
		DeptID:     request.DeptID,
		Level:      request.Level,
		JoinDate:   request.JoinDate,
		IsActive:   false,
		Address:    request.Address,
		Email:      request.Email,
		Phone:      request.Phone,
		DOB:        request.DOB,
		POB:        request.POB,
		IsDeleted:  false,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	data = append(data, row)

	err = m.Repository.Record(data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

func (m *AddMasterEmployeesService) generateID(date time.Time) (int64, error) {
	id, err := m.Repository.FindLastID(date)
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}
