package main

import (
    "log"
    "strings"

    "eulabsmyapp/config/database"
    productController "eulabsmyapp/go_modules/products/controller"
    productRepository "eulabsmyapp/core/products/infrastructure/repository"
    createProduct "eulabsmyapp/core/products/application/usecase/createProduct"
    getProduct "eulabsmyapp/core/products/application/usecase/getProduct"
    updateProduct "eulabsmyapp/core/products/application/usecase/updateProduct"
    deleteProduct "eulabsmyapp/core/products/application/usecase/deleteProduct"
    userController "eulabsmyapp/go_modules/users/controller"
    userRepository "eulabsmyapp/core/users/infrastructure/repository"
    registerUser "eulabsmyapp/core/users/application/usecase/registerUser"
    loginUser "eulabsmyapp/core/users/application/usecase/loginUser"
    orderController "eulabsmyapp/go_modules/orders/controller"
    orderRepository "eulabsmyapp/core/orders/infrastructure/repository"
    createOrder "eulabsmyapp/core/orders/application/usecase/createOrder"
    getOrder "eulabsmyapp/core/orders/application/usecase/getOrder"
    updateOrder "eulabsmyapp/core/orders/application/usecase/updateOrder"
    deleteOrder "eulabsmyapp/core/orders/application/usecase/deleteOrder"
    productValidator "eulabsmyapp/core/products/infrastructure/validator"
    userValidator "eulabsmyapp/core/users/infrastructure/validator"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

var jwtSecret = []byte("supersecretkey")

func main() {
    db, err := database.InitDB()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer db.Close()

    e := echo.New()

    e.Validator = productValidator.NewValidator()
    e.Validator = userValidator.NewValidator()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: jwtSecret,
        Skipper: func(c echo.Context) bool {
            path := c.Path()
            return path == "/login" || path == "/register" || strings.HasPrefix(path, "/health")
        },
    }))

    productRepo := productRepository.NewProductRepository(db)
    productCreateUsecase := createProduct.NewCreateProductUsecase(productRepo)
    productGetUsecase := getProduct.NewGetProductUsecase(productRepo)
    productUpdateUsecase := updateProduct.NewUpdateProductUsecase(productRepo)
    productDeleteUsecase := deleteProduct.NewDeleteProductUsecase(productRepo)

    productCtrl := productController.NewProductController(productCreateUsecase, productGetUsecase, productUpdateUsecase, productDeleteUsecase)

    e.POST("/product/create", productCtrl.CreateProduct)
    e.GET("/product/:id", productCtrl.GetProduct)
    e.PUT("/product/:id", productCtrl.UpdateProduct)
    e.DELETE("/product/:id", productCtrl.DeleteProduct)

    userRepo := userRepository.NewUserRepository(db)
    registerUsecase := registerUser.NewRegisterUserUsecase(userRepo)
    loginUsecase := loginUser.NewLoginUserUsecase(userRepo)
    userCtrl := userController.NewUserController(registerUsecase, loginUsecase)

    e.POST("/register", userCtrl.Register)
    e.POST("/login", userCtrl.Login)

    orderRepo := orderRepository.NewOrderRepository(db)
    orderCreateUsecase := createOrder.NewCreateOrderUsecase(orderRepo)
    orderGetUsecase := getOrder.NewGetOrderUsecase(orderRepo)
    orderUpdateUsecase := updateOrder.NewUpdateOrderUsecase(orderRepo)
    orderDeleteUsecase := deleteOrder.NewDeleteOrderUsecase(orderRepo)

    orderCtrl := orderController.NewOrderController(orderCreateUsecase, orderGetUsecase, orderUpdateUsecase, orderDeleteUsecase)

    e.POST("/order/create", orderCtrl.CreateOrder)
    e.GET("/order/:id", orderCtrl.GetOrder)
    e.PUT("/order/update/:id", orderCtrl.UpdateOrder)
    e.DELETE("/orders/delete/:id", orderCtrl.DeleteOrder)

    e.GET("/health", func(c echo.Context) error {
        return c.String(200, "OK")
    })

    e.Logger.Fatal(e.Start(":8080"))
}
