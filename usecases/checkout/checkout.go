package checkout

import (
	"github.com/wendylau87/bcg/domain/checkout"
	"github.com/wendylau87/bcg/domain/item"
	"github.com/wendylau87/bcg/domain/promotion"
	"github.com/wendylau87/bcg/entities"
	"github.com/wendylau87/bcg/infrastructure/logger"
)

type UsecaseItf interface {
	CreateCheckout(v entities.CreateCheckOut) (*entities.CheckOutResp, error)
}

type usecase struct {
	logger logger.Logger
	itemdom item.DomainItf
	promotiondom promotion.DomainItf
	checkoutdom checkout.DomainItf
}



func InitUsecase(logger logger.Logger, itemdom item.DomainItf, promotiondom promotion.DomainItf, checkoutdom checkout.DomainItf) UsecaseItf {
	return &usecase{
		logger,
		itemdom,
		promotiondom,
		checkoutdom,
	}
}
