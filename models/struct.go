package models

import "gorm.io/gorm"

type Environment struct {
	gorm.Model
	Water int `json:"water" gorm:"not null"`
	Wind  int `json:"wind" gorm:"wind"`
}

type Response struct {
	Messages string
	Success  bool
	Data     interface{}
}
