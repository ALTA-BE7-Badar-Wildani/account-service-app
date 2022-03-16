package entity

import "gorm.io/gorm"

type TopUp struct {
	gorm.Model
	UserId int
	Nominal uint
	User User `gorm:"foreignKey:UserId"`
}

