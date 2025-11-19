package services

import (
	"context"
	"database/sql"

	"github.com/ojipoetra/backend-web-tonggak/helper"
	"github.com/ojipoetra/backend-web-tonggak/models/domain"
	"github.com/ojipoetra/backend-web-tonggak/models/web"
	"github.com/ojipoetra/backend-web-tonggak/repository"
)

type KamarServiceImpl struct {
	KamarRepository repository.KamarRepository
	DB              *sql.DB
}

// Create implements KamarService.
func (service *KamarServiceImpl) Create(ctx context.Context, request web.CreateKamarRequest) web.KamarResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func(){
		err := recover()
		if err != nil{
			errorRolback := tx.Rollback()
			helper.PanicError(errorRolback)
			panic(err)
		}else{
			errorCommit := tx.Commit()
			helper.PanicError(errorCommit)
		}
	}()

	kamar := domain.Kamar{
		NamaKamar: request.NamaKamar,
		Terpakai: request.Terpakai,
		Tersedia: request.Tersedia,
	}
	kamar = service.KamarRepository.Save(ctx, tx, kamar)
	return helper.ToCreateResponse(kamar)
}


// Delete implements KamarService.
func (service *KamarServiceImpl) Delete(ctx context.Context, kamarId int) {
	panic("unimplemented")
}

// FindAll implements KamarService.
func (service *KamarServiceImpl) FindAll(ctx context.Context) []web.KamarResponse {
	panic("unimplemented")
}

// FindById implements KamarService.
func (service *KamarServiceImpl) FindById(ctx context.Context, kamarId int) web.KamarResponse {
	panic("unimplemented")
}

// Update implements KamarService.
func (service *KamarServiceImpl) Update(ctx context.Context, request web.CreateKamarRequest) web.KamarResponse {
	panic("unimplemented")
}

func NewKamarService(KamarRepository repository.KamarRepository, DB *sql.DB) KamarService {
	return &KamarServiceImpl{
		KamarRepository: KamarRepository,
		DB:              DB,
	}
}
