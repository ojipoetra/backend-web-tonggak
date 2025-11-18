package domain

import (
	"time"
)

type Kamar struct{
	Id int `gorm:"PrimaryKey" json:"id"`
	NamaKamar string `gorm:"type:varchar(200)" json:"nama_kamar"`
	Terpakai string `gorm:"type:varchar(10)" json:"terpakai"`
	Tersedia string `gorm:"type:varchar(10)" json:"tersedia"`
	Update time.Time
}