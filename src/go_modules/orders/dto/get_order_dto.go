package dto

type GetOrderDTO struct {
    ID string `json:"id" validate:"required"`
}
