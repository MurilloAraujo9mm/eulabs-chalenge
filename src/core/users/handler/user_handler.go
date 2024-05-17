package handler

import (
	"net/http"

	"myapp/core/users/model"
	"myapp/core/users/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	if user.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username is required"})
	}
	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Password is required"})
	}

	err := h.UserUsecase.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	if user.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Username is required"})
	}
	if user.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Password is required"})
	}

	token, err := h.UserUsecase.Login(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
