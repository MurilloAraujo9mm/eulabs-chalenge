package controller

import (
    "myapp/go_modules/users/dto"
    registerUser "myapp/core/users/application/usecase/registerUser"
    loginUser "myapp/core/users/application/usecase/loginUser"
    "net/http"

    "github.com/labstack/echo/v4"
)

type UserController struct {
    RegisterUsecase registerUser.RegisterUserUsecase
    LoginUsecase    loginUser.LoginUserUsecase
}

func NewUserController(registerUsecase registerUser.RegisterUserUsecase, loginUsecase loginUser.LoginUserUsecase) *UserController {
    return &UserController{
        RegisterUsecase: registerUsecase,
        LoginUsecase:    loginUsecase,
    }
}

func (ctrl *UserController) Register(c echo.Context) error {
    var registerDTO dto.RegisterDTO
    if err := c.Bind(&registerDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
    }
    if err := c.Validate(&registerDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    user, err := ctrl.RegisterUsecase.Execute(registerDTO)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}

func (ctrl *UserController) Login(c echo.Context) error {
    var loginDTO dto.LoginDTO
    if err := c.Bind(&loginDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
    }
    if err := c.Validate(&loginDTO); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    token, err := ctrl.LoginUsecase.Execute(loginDTO.Username, loginDTO.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
    }

    return c.JSON(http.StatusOK, echo.Map{"token": token})
}
