package models

import "gorm.io/gorm"

type Identifier struct {
	gorm.Model
	ID int `json:"id" gorm:"autoIncrement"`
}
