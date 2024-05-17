package controller

import (
    "eulabsmyapp/go_modules/products/dto"
    createProduct "eulabsmyapp/core/products/application/usecase/createProduct"
    getProduct "eulabsmyapp/core/products/application/usecase/getProduct"
    updateProduct "eulabsmyapp/core/products/application/usecase/updateProduct"
    deleteProduct "eulabsmyapp/core/products/application/usecase/deleteProduct"
    "net/http"

    "github.com/labstack/echo/v4"
)

type ProductController struct {
    CreateUsecase createProduct.CreateProductUsecase
    GetUsecase    getProduct.GetProductUsecase
    UpdateUsecase updateProduct.UpdateProductUsecase
    DeleteUsecase deleteProduct.DeleteProductUsecase
}

func NewProductController(
    createUsecase createProduct.CreateProductUsecase,
    getUsecase getProduct.GetProductUsecase,
    updateUsecase updateProduct.UpdateProductUsecase,
    deleteUsecase deleteProduct.DeleteProductUsecase,
) *ProductController {
    return &ProductController{
        CreateUsecase: createUsecase,
        GetUsecase:    getUsecase,
        UpdateUsecase: updateUsecase,
        DeleteUsecase: deleteUsecase,
    }
}

func (ctrl *ProductController) CreateProduct(c echo.Context) error {
    var createProductDTO dto.CreateProductDTO
    if err := c.Bind(&createProductDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
    }
    if err := c.Validate(&createProductDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    product, err := ctrl.CreateUsecase.Execute(createProductDTO)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, product)
}

func (ctrl *ProductController) GetProduct(c echo.Context) error {
    id := c.Param("id")

    product, err := ctrl.GetUsecase.Execute(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
    }

    return c.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) UpdateProduct(c echo.Context) error {
    id := c.Param("id")

    var updateProductDTO dto.UpdateProductDTO
    if err := c.Bind(&updateProductDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
    }
    if err := c.Validate(&updateProductDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    product, err := ctrl.UpdateUsecase.Execute(id, updateProductDTO)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, product)
}

func (ctrl *ProductController) DeleteProduct(c echo.Context) error {
    id := c.Param("id")

    err := ctrl.DeleteUsecase.Execute(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "Product deleted successfully"})
}
