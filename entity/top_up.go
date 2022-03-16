package entity

import "gorm.io/gorm"

type TopUp struct {
	gorm.Model
	UserId uint
	Nominal uint
	User User `gorm:"foreignKey:UserId"`
}

func (topUp TopUp) TableName() string {
	return "top_up"
}

