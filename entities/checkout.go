package entities

import "github.com/wendylau87/bcg/infrastructure/sqlx"

type CheckOut struct {
	ID           int              `json:"id"`
	CustomerName string           `json:"customer_name"`
	TotalPrice   float64          `json:"total_price"`
	CreatedAt    sqlx.NullTime    `json:"created_at"`
	Details      []CheckOutDetail `json:"details"`
}

type CheckOutDetail struct {
	ID              int     `json:"id"`
	CheckOutID      int     `json:"check_out_id"`
	ItemID          int     `json:"item_id"`
	PromotionItemID int     `json:"promotion_item_id"`
	TotalQuantity   int     `json:"total_quantity"`
	SubTotalPrice   float64 `json:"sub_total_price"`
}

type CreateCheckOut struct {
	CustomerName string                 `json:"customer_name"`
	Details      []CreateCheckOutDetail `json:"details"`
}

type CreateCheckOutDetail struct {
	SKU           string `json:"sku"`
	TotalQuantity int    `json:"total_quantity"`
}

type CheckOutResp struct {
	ID           int                  `json:"id"`
	CustomerName string               `json:"customer_name"`
	TotalPrice   float64              `json:"total_price"`
	CreatedAt    sqlx.NullTime        `json:"created_at"`
	Details      []CheckOutDetailResp `json:"details"`
}

type CheckOutDetailResp struct {
	ItemName      string  `json:"item_name"`
	PromotionName string  `json:"promotion_name"`
	TotalQuantity int     `json:"total_quantity"`
	SubTotalPrice float64 `json:"sub_total_price"`
}
