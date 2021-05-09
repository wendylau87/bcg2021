package checkout

import (
	"github.com/wendylau87/bcg/entities"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/infrastructure/sqlhandler"
)

type DomainItf interface {
	CreateCheckout(v entities.CheckOut)(*entities.CheckOut, error)
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
