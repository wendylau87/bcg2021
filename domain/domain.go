package domain

import (
	"github.com/wendylau87/bcg/domain/checkout"
	"github.com/wendylau87/bcg/domain/item"
	"github.com/wendylau87/bcg/domain/promotion"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/infrastructure/sqlhandler"
)

type Domain struct {
	Item      item.DomainItf
	Promotion promotion.DomainItf
	Checkout  checkout.DomainItf
}

func Init(
	logger logger.Logger,
	sql sqlhandler.SQLHandler,
) *Domain {
	return &Domain{
		Item : item.InitDomain(logger, sql),
		Promotion : promotion.InitDomain(logger, sql),
		Checkout : checkout.InitDomain(logger, sql),
	}
}