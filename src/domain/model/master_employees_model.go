package model

type MasterEmployeesResponseModel struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	DeptID   string `json:"dept_id" db:"dept_id"`
	Level    string `json:"level" db:"level"`
	JoinDate string `json:"join_date" db:"join_date"`
	IsActive bool   `json:"is_active" db:"is_active"`
	Address  string `json:"address" db:"address"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	DOB      string `json:"dob" db:"dob"`
	POB      string `json:"pob" db:"pob"`
}

type MasterEmployeesRequestModel struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	DeptID   string `json:"dept_id" db:"dept_id"`
	Level    string `json:"level" db:"level"`
	JoinDate string `json:"join_date" db:"join_date"`
	Address  string `json:"address" db:"address"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	DOB      string `json:"dob" db:"dob"`
	POB      string `json:"pob" db:"pob"`
}

type MasterEmployeesActivateRequestModel struct {
	ID       int64 `json:"id" db:"id"`
	IsActive bool  `json:"is_active" db:"is_active"`
}
