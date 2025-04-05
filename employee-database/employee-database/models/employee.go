package models

type Employee struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Position  string `json:"position"`
	Salary    int    `json:"salary"`
}
