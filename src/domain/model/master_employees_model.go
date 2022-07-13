package model

type MasterEmployeesResponseModel struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	DeptID       string `json:"dept_id" db:"dept_id"`
	Level        string `json:"level" db:"level"`
	JoinDate     string `json:"join_date" db:"join_date"`
	IsActive     int    `json:"is_active" db:"is_active"`
	Address      string `json:"address" db:"address"`
	Email        string `json:"email" db:"email"`
	Phone        string `json:"phone" db:"phone"`
	DOB          string `json:"dob" db:"dob"`
	PlaceOfBirth string `json:"place_of_birth" db:"place_of_birth"`
}
