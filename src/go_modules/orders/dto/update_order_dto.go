package dto

type UpdateOrderDTO struct {
    Quantity int     `json:"quantity" validate:"required"`
    Total    float64 `json:"total" validate:"required"`
}
