package models

import "time"

// OrderItem ..
type OrderItem struct {
	ID          string     `json:"id"`
	ProductName string     `json:"productName" binding:"required" example:"Lavash"`
	Client_id   string     `json:"client_id" binding:"required" example:"uuid"`
	TotalPrice  string     `json:"totalPrice" binding:"required" example:"uuid"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// CreateOrderItemModul ..
type CreateOrderItemModul struct {
	ProductName string `json:"productName" binding:"required" example:"Lavash"`
	Client_id   string `json:"client_id" binding:"required" example:"uuid"`
	TotalPrice  string `json:"totalPrice" binding:"required" example:"uuid"`
}
