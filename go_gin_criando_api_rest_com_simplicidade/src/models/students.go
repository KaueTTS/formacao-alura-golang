package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
	RG    string `json:"rg"`
}

var Students []Student
