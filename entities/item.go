package entities

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Sku      string  `json:"sku"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}
