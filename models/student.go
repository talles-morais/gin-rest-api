package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Phone string `json:"phone"`
}

var Students []Student
