package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama string
	JenisKelamin string
	Alamat string
	NomorHP string
	Saldo uint
	Email string
	TanggalLahir time.Time
	TopUp []TopUp `gorm:"foreignKey:UserId;references:ID"`
}
