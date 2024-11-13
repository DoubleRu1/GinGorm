package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
