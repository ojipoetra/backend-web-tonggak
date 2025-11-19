package helper

import (
	"github.com/ojipoetra/backend-web-tonggak/models/domain"
	"github.com/ojipoetra/backend-web-tonggak/models/web"
)
func ToCreateResponse (kamar domain.Kamar) web.KamarResponse{
	return web.KamarResponse{
		Id: kamar.Id,
		NamaKamar: kamar.NamaKamar,
		Terpakai: kamar.Terpakai,
		Tersedia: kamar.Tersedia,
	}
}