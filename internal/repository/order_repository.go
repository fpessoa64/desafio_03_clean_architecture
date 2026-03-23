package repository

import (
	"context"

	"github.com/fpessoa64/desafio03_clean_arch/internal/entity"
)

type OrderRepository interface {
	Create(ctx context.Context, o *entity.Order) error
	List(ctx context.Context) ([]entity.Order, error)
}
