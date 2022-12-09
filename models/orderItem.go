package models

import "time"

// OrderItem ..
type OrderItem struct {
	ID         string     `json:"id"`
	Client_id  string     `json:"client_id" binding:"required" example:"uuid"`
	TotalPrice string     `json:"totalPrice" binding:"required" example:"uuid"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

// CreateOrderItemModul ..
type CreateOrderItemModul struct {
	Client_id  string `json:"client_id" binding:"required" example:"uuid"`
	TotalPrice string `json:"totalPrice" binding:"required" example:"uuid"`
}
