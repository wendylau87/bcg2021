package checkout

import "github.com/wendylau87/bcg/entities"

func (d *domain) CreateCheckout(v entities.CheckOut) (*entities.CheckOut, error) {
	return d.createCheckout(v)
}