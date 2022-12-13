package models

import "time"

// OrderItem ..
type OrderItem struct {
	ID         string     `json:"id"`
	OrderId    string     `json:"orderId" binding:"required" example:"111"`
	TotalPrice int32      `json:"totalPrice" binding:"required" example:"111"`
	Count      int32      `json:"count" binding:"required" example:"111"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

// CreateOrderItemModul ..
type CreateOrderItemModul struct {
	OrderId    string `json:"orderId" binding:"required" example:"111"`
	TotalPrice int32  `json:"totalPrice" binding:"required" example:"111"`
}
