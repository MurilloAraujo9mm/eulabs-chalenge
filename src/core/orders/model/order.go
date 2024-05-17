package model

import "myapp/core/products/model"

type Order struct {
	ID       int          `json:"id"`
	UserID   int          `json:"user_id"`
	Products []model.Product `json:"products"`
	Quantity int          `json:"quantity"`
}
