package entity

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	UserId int
	UserPenerimaId int
	Nominal uint
	User User `gorm:"foreignKey:UserId"`
	UserPenerima User `gorm:"foreignKey:UserId"`
}
