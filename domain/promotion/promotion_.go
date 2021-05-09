package promotion

import "github.com/wendylau87/bcg/entities"

func (d *domain) GetPromotionByID(id int) (*entities.Promotion, error) {
	return d.findPromotionByID(id)
}

func (d *domain) GetPromotionItemByID(id int) (*entities.PromotionItem, error) {
	return d.findPromotionItemByID(id)
}
func (d *domain)GetAllPromotionItemByItemID(id int)([]entities.PromotionItem, error){
	return d.findAllPromotionItemByItemID(id)
}
