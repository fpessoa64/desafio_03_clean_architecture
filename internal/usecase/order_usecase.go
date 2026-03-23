package usecase

import (
	"context"

	"github.com/fpessoa64/desafio03_clean_arch/internal/entity"
	"github.com/fpessoa64/desafio03_clean_arch/internal/repository"
)

type OrderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) Create(ctx context.Context, o *entity.Order) error {
	return u.repo.Create(ctx, o)
}

func (u *OrderUsecase) List(ctx context.Context) ([]entity.Order, error) {
	return u.repo.List(ctx)
}
