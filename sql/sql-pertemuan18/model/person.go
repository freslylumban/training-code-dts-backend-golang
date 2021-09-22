package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	ID    uint
	Name  string
	Email string
}
