package handler

import (
	"net/http"
	"strconv"

	"myapp/core/orders/model"
	"myapp/core/orders/usecase"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderUsecase usecase.OrderUsecase
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	order := new(model.Order)
	if err := c.Bind(order); err != nil {
		return err
	}

	err := h.OrderUsecase.CreateOrder(order)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.OrderUsecase.GetOrder(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Order not found"})
	}

	return c.JSON(http.StatusOK, order)
}
