package services

import (
	"context"

	"github.com/ojipoetra/backend-web-tonggak/models/web"
)

type KamarService interface{
	Create(ctx context.Context, request web.CreateKamarRequest) web.KamarResponse
	Update(ctx context.Context, request web.CreateKamarRequest) web.KamarResponse
	FindById(ctx context.Context, kamarId int)web.KamarResponse
	Delete(ctx context.Context, kamarId int)
	FindAll(ctx context.Context)[]web.KamarResponse
}