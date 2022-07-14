package employees

import (
	"fmt"
	"net/http"
	"simple-go-app/src/domain/exception"
	"simple-go-app/src/domain/interfaces"
	"simple-go-app/src/domain/model"
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

func (m *EditMasterEmployeesService) Edit(request *model.MasterEmployeesModel) (response int, err error) {
	var data []*model.MasterEmployeesModel
	var isExist bool

	isExist, err = m.Repository.IfExistByID(fmt.Sprintf("%v", request.ID))
	exception.PanicIfNeeded(err)

	if !isExist {
		return http.StatusNotFound, nil
	}

	data = append(data, request)

	err = m.Repository.Insert(data)
	exception.PanicIfNeeded(err)

	return http.StatusOK, nil
}
