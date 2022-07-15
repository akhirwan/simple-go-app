package employees

import (
	"net/http"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/interfaces"
)

type RemoveMasterEmployeesService struct {
	Repository interfaces.MasterEmployeesRepositoryInterfaces
}

func NewRemoveMasterEmployeesService(
	repository interfaces.MasterEmployeesRepositoryInterfaces,
) *RemoveMasterEmployeesService {
	return &RemoveMasterEmployeesService{
		Repository: repository,
	}
}

func (m *RemoveMasterEmployeesService) Remove(id string) (int, error) {
	err := m.Repository.Delete(id)
	exception.PanicIfNeeded(err)

	return http.StatusOK, nil
}
