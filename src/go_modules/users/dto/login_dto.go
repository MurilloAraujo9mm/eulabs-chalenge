package dto

type LoginDTO struct {
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}
