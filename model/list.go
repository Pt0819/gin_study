package model

import "gorm.io/gorm"

type list struct {
	gorm.Model
	ID    uint   `json:"ID" gorm:"id"`
	Name  string `json:"Name" gorm:"id"`
	Phone uint   `json:"Phone" gorm:"phone"`
}
