package main

import (
    "log"
    productController "myapp/core/products/controller"
    productRepository "myapp/core/products/repository"
    createProduct "myapp/core/products/usecase/createProduct"
    getProduct "myapp/core/products/usecase/getProduct"
    updateProduct "myapp/core/products/usecase/updateProduct"
    deleteProduct "myapp/core/products/usecase/deleteProduct"
    userController "myapp/core/users/controller"
    userRepository "myapp/core/users/repository"
    registerUser "myapp/core/users/usecase/registerUser"
    loginUser "myapp/core/users/usecase/loginUser"
    "myapp/infrastructure/database"
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

    e.Logger.Fatal(e.Start(":8080"))
}
