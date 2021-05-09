package entities

type Promotion struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	MinimumItem         int64   `json:"minimum_item"`
	MinimumItemDiscount float64 `json:"minimum_item_discount"`
	CountOrigin         int64   `json:"count_origin"`
	CountDestination    int64   `json:"count_destination"`
	FreeItem            bool    `json:"free_item"`
}

type PromotionItem struct {
	ID          int  `json:"id"`
	PromotionID int  `json:"promotion_id"`
	ItemID      int  `json:"item_id"`
	FreeItemID  int  `json:"free_item_id"`
	IsActive    bool `json:"is_active"`
	Priority    int  `json:"priority"`
}
