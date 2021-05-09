package item

import (
	"github.com/wendylau87/bcg/entities"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/infrastructure/sqlhandler"
)

type DomainItf interface {
	GetItemByID(id int)(*entities.Item, error)
	GetItemBySku(sku string)(*entities.Item, error)
}

type domain struct {
	logger logger.Logger
	SQLHandler sqlhandler.SQLHandler
}

func InitDomain(logger logger.Logger, sql sqlhandler.SQLHandler) DomainItf {
	return &domain{
		logger,
		sql,
	}
}