package web

import "time"

type CreateKamarRequest struct{
	NamaKamar string
	Terpakai string
	Tersedia string
	Update time.Time
}