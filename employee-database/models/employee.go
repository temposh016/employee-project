package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string  `json:"name"`
	Position   string  `json:"position"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}
