package usecases

import (
	"github.com/wendylau87/bcg/domain"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/usecases/checkout"
)

type Usecase struct {
	Checkout checkout.UsecaseItf
}

func Init(logger logger.Logger, dom *domain.Domain) *Usecase {
	return &Usecase{
		Checkout : checkout.InitUsecase(logger, dom.Item, dom.Promotion, dom.Checkout),
	}
}