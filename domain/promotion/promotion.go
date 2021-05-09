package promotion

import (
	"github.com/wendylau87/bcg/entities"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/infrastructure/sqlhandler"
)

type DomainItf interface {
	GetPromotionByID(id int)(*entities.Promotion, error)
	GetPromotionItemByID(id int)(*entities.PromotionItem, error)
	GetAllPromotionItemByItemID(id int)([]entities.PromotionItem, error)
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
