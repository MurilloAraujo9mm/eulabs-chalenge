package controller

import (
    "myapp/go_modules/orders/dto"
    createOrder "myapp/core/orders/application/usecase/createOrder"
    getOrder "myapp/core/orders/application/usecase/getOrder"
    updateOrder "myapp/core/orders/application/usecase/updateOrder"
    deleteOrder "myapp/core/orders/application/usecase/deleteOrder"
    "net/http"

    "github.com/labstack/echo/v4"
)

type OrderController struct {
    CreateUsecase createOrder.CreateOrderUsecase
    GetUsecase    getOrder.GetOrderUsecase
    UpdateUsecase updateOrder.UpdateOrderUsecase
    DeleteUsecase deleteOrder.DeleteOrderUsecase
}

func NewOrderController(
    createUsecase createOrder.CreateOrderUsecase,
    getUsecase getOrder.GetOrderUsecase,
    updateUsecase updateOrder.UpdateOrderUsecase,
    deleteUsecase deleteOrder.DeleteOrderUsecase,
) *OrderController {
    return &OrderController{
        CreateUsecase: createUsecase,
        GetUsecase:    getUsecase,
        UpdateUsecase: updateUsecase,
        DeleteUsecase: deleteUsecase,
    }
}

func (ctrl *OrderController) CreateOrder(c echo.Context) error {
    var createOrderDTO dto.CreateOrderDTO
    if err := c.Bind(&createOrderDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
    }
    if err := c.Validate(&createOrderDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    order, err := ctrl.CreateUsecase.Execute(createOrderDTO)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, order)
}

func (ctrl *OrderController) GetOrder(c echo.Context) error {
    id := c.Param("id")

    order, err := ctrl.GetUsecase.Execute(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "Order not found"})
    }

    return c.JSON(http.StatusOK, order)
}

func (ctrl *OrderController) UpdateOrder(c echo.Context) error {
    id := c.Param("id")

    var updateOrderDTO dto.UpdateOrderDTO
    if err := c.Bind(&updateOrderDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
    }
    if err := c.Validate(&updateOrderDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    order, err := ctrl.UpdateUsecase.Execute(id, updateOrderDTO)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, order)
}

func (ctrl *OrderController) DeleteOrder(c echo.Context) error {
    id := c.Param("id")

    err := ctrl.DeleteUsecase.Execute(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "Order deleted successfully"})
}
