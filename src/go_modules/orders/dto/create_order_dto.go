package dto

type CreateOrderDTO struct {
    UserID    string  `json:"user_id" validate:"required"`
    ProductID string  `json:"product_id" validate:"required"`
    Quantity  int     `json:"quantity" validate:"required"`
    Total     float64 `json:"total" validate:"required"`
}
