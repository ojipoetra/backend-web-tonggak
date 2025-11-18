package repository

import (
	"context"
	"database/sql"

	"github.com/ojipoetra/backend-web-tonggak/models/domain"
)

type KamarRepository interface{
	Save(ctx context.Context, tx *sql.Tx, kamar domain.Kamar)domain.Kamar
	Update(ctx context.Context, tx *sql.Tx, kamar domain.Kamar)domain.Kamar
	Delete(ctx context.Context, tx *sql.Tx, kamar domain.Kamar)domain.Kamar
	FindById(ctx context.Context, tx *sql.Tx, kamar domain.Kamar)domain.Kamar
	FindAll(ctx context.Context, tx *sql.Tx)[]domain.Kamar
}

