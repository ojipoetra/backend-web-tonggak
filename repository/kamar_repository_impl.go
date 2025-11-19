package repository

import (
	"context"
	"database/sql"

	"github.com/ojipoetra/backend-web-tonggak/models/domain"
)

type KamarRepositoryImpl struct {
}

// Delete implements KamarRepository.
func (repository *KamarRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, kamar domain.Kamar) domain.Kamar {
	// SQL := "insert into kamars(nama_kamar,terpakai,tersedia,update) values (?,?,?,?)"
	// result, err := tx.ExecContext(ctx, SQL, kamar.NamaKamar,kamar.Terpakai,kamar.Tersedia,kamar.Update)
	// if err != nil{
	// 	panic(err)
	// }
	// kamar.Id = int
	panic("unimplemented")
}

// FindAll implements KamarRepository.
func (repository *KamarRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Kamar {
	panic("unimplemented")
}

// FindById implements KamarRepository.
func (repository *KamarRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, kamar domain.Kamar) domain.Kamar {
	panic("unimplemented")
}

// Save implements KamarRepository.
func (repository *KamarRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, kamar domain.Kamar) domain.Kamar {
	SQL := "insert into kamars(nama_kamar, terpakai, tersedia) values (?,?,?)"
	result, err := tx.ExecContext(ctx,SQL,&kamar.NamaKamar,&kamar.Terpakai,&kamar.Tersedia)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	kamar.Id = int(id)
	return kamar
}

// Update implements KamarRepository.
func (repository *KamarRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, kamar domain.Kamar) domain.Kamar {
	SQL := "update kamars set nama_kamar, terpakai, tersedia, update where id = ?"
	_, err := tx.ExecContext(ctx, SQL, kamar.Id)
	if err != nil {
		panic(err)
	}
	return kamar
}

func NewKamarRepository() KamarRepository {
	return &KamarRepositoryImpl{}
}
