package controllers

import (
	"encoding/json"
	"github.com/wendylau87/bcg/entities"
	"github.com/wendylau87/bcg/infrastructure/logger"
	"github.com/wendylau87/bcg/usecases"
	"net/http"
)

type CheckoutController struct {
	Usecase *usecases.Usecase
	Logger  logger.Logger
}

func InitCheckoutController(uc *usecases.Usecase, logger logger.Logger) *CheckoutController {
	return &CheckoutController{
		Usecase: uc,
		Logger: logger,
	}
}

func (c *CheckoutController) CreateCheckout(w http.ResponseWriter, r *http.Request) {
	c.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	var createCheckout entities.CreateCheckOut
	err := json.NewDecoder(r.Body).Decode(&createCheckout)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	resp, err := c.Usecase.Checkout.CreateCheckout(createCheckout)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResult, err := json.Marshal(resp)
	if err != nil {
		c.Logger.LogError("%s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Write(jsonResult)
}
