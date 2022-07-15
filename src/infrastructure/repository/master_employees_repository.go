package repository

import (
	"simple-go-app/src/infrastructure/entity"
	"strconv"
	"time"

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

func (m *MasterEmployeesRepository) Record(data []*entity.MasterEmployeesEntity) error {
	tx := m.DB.MustBegin()

	_, err := tx.NamedExec(
		`INSERT INTO master_employee (
			id,
			name,
			dept_id,
			level,
			join_date,
			is_active,
			address,
			email,
			phone,
			dob,
			pob,
			is_deleted,
			created_at,
			modified_at
		) VALUES (
			:id,
			:name,
			:dept_id,
			:level,
			:join_date,
			:is_active,
			:address,
			:email,
			:phone,
			:dob,
			:pob,
			:is_deleted,
			:created_at,
			:modified_at
		) ON DUPLICATE KEY UPDATE
		name		= VALUES(name),
		level		= VALUES(level),
		join_date	= VALUES(join_date),
		is_active	= VALUES(is_active),
		address		= VALUES(address),
		email		= VALUES(email),
		phone		= VALUES(phone),
		dob			= VALUES(dob),
		pob			= VALUES(pob),
		is_deleted	= VALUES(is_deleted),
		modified_at	= VALUES(modified_at);`,
		data)

	if err != nil {
		pp.Println("[FATAL] From insert master_employees : ", err)
		tx.Rollback()
		return err
	} else {
		err = tx.Commit()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MasterEmployeesRepository) FindAll() (data []*entity.MasterEmployeesEntity, err error) {

	tx := m.DB.MustBegin()

	err = tx.Select(&data, `SELECT * FROM master_employee WHERE is_deleted = 0 ORDER BY id DESC;;`)

	if err != nil {
		pp.Println("[FATAL] From read master_employees : ", err)
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

func (m *MasterEmployeesRepository) FindByID(id string) (*entity.MasterEmployeesEntity, error) {
	var data []*entity.MasterEmployeesEntity

	tx := m.DB.MustBegin()

	err := tx.Select(&data, `SELECT * FROM master_employee WHERE id = ? AND is_deleted = 0;`, id)

	if err != nil {
		pp.Println("[FATAL] From read master_employees : ", err)
		tx.Rollback()
		return nil, err
	} else {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	if len(data) == 0 {
		return nil, err
	}

	return data[0], nil
}

func (m *MasterEmployeesRepository) FindExistingByID(id string) (*entity.MasterEmployeesEntity, error) {
	var data []*entity.MasterEmployeesEntity

	tx := m.DB.MustBegin()

	err := tx.Select(&data, `SELECT * FROM master_employee WHERE id = ?;`, id)

	if err != nil {
		pp.Println("[FATAL] From read master_employees : ", err)
		tx.Rollback()
		return nil, err
	} else {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	if len(data) == 0 {
		return nil, err
	}

	return data[0], nil
}

func (m *MasterEmployeesRepository) FindLastID(date time.Time) (result int64, err error) {
	var data []*entity.MasterEmployeesEntity

	tx := m.DB.MustBegin()

	err = tx.Select(&data,
		`SELECT IFNULL(MAX(id), 0) id FROM master_employee WHERE LEFT(id, 8) = ?;`,
		date.Format("20060102"))

	if err != nil {
		pp.Println("[FATAL] From read master_employees : ", err.Error())
		tx.Rollback()
		return 0, err
	} else {
		err = tx.Commit()
		if err != nil {
			return 0, err
		}
	}

	result = data[0].ID
	if data[0].ID == 0 {
		result, _ = strconv.ParseInt(date.Format("20060102")+"000", 10, 64)
	}

	return result, nil
}

func (m *MasterEmployeesRepository) Delete(id string) error {
	tx := m.DB.MustBegin()

	_, err := tx.NamedExec(
		`DELETE FROM master_employee WHERE id = :id;`,
		map[string]interface{}{"id": id})

	if err != nil {
		pp.Println("[FATAL] From read master_employees : ", err.Error())
		tx.Rollback()
		return err
	} else {
		err = tx.Commit()
		if err != nil {
			return err
		}
	}

	return nil
}
