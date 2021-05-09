package main

import (
	"github.com/wendylau87/bcg/infrastructure"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/infrastructure/sqlhandler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := logger.NewLogger()

	infrastructure.Load(*logger)
	logger.LogAccess("Logger initialized...")

	sqlHandler, err := sqlhandler.NewSQLHandler(*logger)
	if err != nil {
		logger.LogError("%s", err)
		panic(err)
	}
	infrastructure.Dispatch(*logger, sqlHandler)


}
