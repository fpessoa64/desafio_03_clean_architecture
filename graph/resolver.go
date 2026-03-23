package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

import "github.com/fpessoa64/desafio03_clean_arch/internal/usecase"

type Resolver struct {
	UC *usecase.OrderUsecase
}
