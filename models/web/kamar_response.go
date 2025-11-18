package web

import "time"

type KamarResponse struct{
	Id int
	NamaKamar string
	Terpakai string
	Tersedia string
	Update time.Time
}