package item

import "github.com/wendylau87/bcg/entities"

func (d *domain) GetItemByID(id int) (*entities.Item, error) {
	return d.findItemByID(id)
}

func (d *domain) GetItemBySku(sku string) (*entities.Item, error) {
	return d.findItemBySKU(sku)
}
