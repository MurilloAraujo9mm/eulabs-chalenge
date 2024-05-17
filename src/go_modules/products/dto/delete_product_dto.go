package dto

type DeleteProductDTO struct {
    ID string `json:"id" validate:"required"`
}
