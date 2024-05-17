package handler

import (
	"net/http"
	"myapp/core/products/model"
	"myapp/core/products/usecase"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)

	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	err := h.ProductUsecase.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")
	product, err := h.ProductUsecase.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	err := h.ProductUsecase.UpdateProduct(id, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := h.ProductUsecase.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Product deleted successfully"})
}
