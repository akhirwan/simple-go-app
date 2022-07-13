package repository

import (
	"simple-go-app/src/domain/model"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

type MasterEmployeesRepository struct {
	DB *sqlx.DB
}

func NewMasterEmployeesRepository(db *sqlx.DB) *MasterEmployeesRepository {
	return &MasterEmployeesRepository{
		DB: db,
	}
}

func (m *MasterEmployeesRepository) FindAll() ([]*model.MasterEmployeesResponseModel, error) {
	var data []*model.MasterEmployeesResponseModel

	tx := m.DB.MustBegin()

	err := tx.Select(
		&data,
		`SELECT * FROM master_employee`)

	if err != nil {
		pp.Println("[FATAL] From : master_employee", err)
		tx.Rollback()
		return nil, err
	} else {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
