package models

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Id     string `json:"id" gorm:"primaryKey"`
	Floor  string `json:"floor"`
	Booked string `json:"booked"`
	From   string `json:"from"`
	To     string `json:"to"`
}
