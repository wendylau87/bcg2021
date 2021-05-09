package infrastructure

import (
	"github.com/wendylau87/bcg/controllers"
	"github.com/wendylau87/bcg/domain"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/infrastructure/sqlhandler"
	"github.com/wendylau87/bcg/usecases"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// Dispatch is handle routing
func Dispatch(logger logger.Logger, sqlHandler sqlhandler.SQLHandler) {
	dom := domain.Init(logger, sqlHandler)
	uc := usecases.Init(logger, dom)
	checkoutController := controllers.InitCheckoutController(uc,logger)

	r := chi.NewRouter()
	r.Post("/checkout", checkoutController.CreateCheckout)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.LogError("%s", err)
	}

	logger.LogAccess("HTTP served on %s", os.Getenv("SERVER_PORT"))
}
