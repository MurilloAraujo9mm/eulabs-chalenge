package dto

type CreateProductDTO struct {
    Name  string  `json:"name" validate:"required"`
    Price float64 `json:"price" validate:"required"`
}
