package main

import (
    "log"
    productController "myapp/go_modules/products/controller"
    productRepository "myapp/core/products/infrastructure/repository"
    createProduct "myapp/core/products/application/usecase/createProduct"
    getProduct "myapp/core/products/application/usecase/getProduct"
    updateProduct "myapp/core/products/application/usecase/updateProduct"
    deleteProduct "myapp/core/products/application/usecase/deleteProduct"
    userController "myapp/go_modules/users/controller"
    userRepository "myapp/core/users/infrastructure/repository"
    registerUser "myapp/core/users/application/usecase/registerUser"
    loginUser "myapp/core/users/application/usecase/loginUser"
    orderController "myapp/go_modules/orders/controller"
    orderRepository "myapp/core/orders/infrastructure/repository"
    createOrder "myapp/core/orders/application/usecase/createOrder"
    getOrder "myapp/core/orders/application/usecase/getOrder"
    updateOrder "myapp/core/orders/application/usecase/updateOrder"
    deleteOrder "myapp/core/orders/application/usecase/deleteOrder"
    "myapp/infrastructure/database"
    productValidator "myapp/core/products/infrastructure/validator"
    userValidator "myapp/core/users/infrastructure/validator"
    "strings"

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
            if c.Path() == "/login" || c.Path() == "/register" || strings.HasPrefix(c.Path(), "/health") {
                return true
            }
            return false
        },
    }))

    productRepo := productRepository.NewProductRepository(db)
    productCreateUsecase := createProduct.NewCreateProductUsecase(productRepo)
    productGetUsecase := getProduct.NewGetProductUsecase(productRepo)
    productUpdateUsecase := updateProduct.NewUpdateProductUsecase(productRepo)
    productDeleteUsecase := deleteProduct.NewDeleteProductUsecase(productRepo)

    productCtrl := productController.NewProductController(productCreateUsecase, productGetUsecase, productUpdateUsecase, productDeleteUsecase)

    e.POST("/products", productCtrl.CreateProduct)
    e.GET("/products/:id", productCtrl.GetProduct)
    e.PUT("/products/:id", productCtrl.UpdateProduct)
    e.DELETE("/products/:id", productCtrl.DeleteProduct)

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

    e.POST("/orders", orderCtrl.CreateOrder)
    e.GET("/orders/:id", orderCtrl.GetOrder)
    e.PUT("/orders/:id", orderCtrl.UpdateOrder)
    e.DELETE("/orders/:id", orderCtrl.DeleteOrder)

    e.Logger.Fatal(e.Start(":8080"))
}
