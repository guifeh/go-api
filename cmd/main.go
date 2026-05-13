package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/middleware"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUseCase(UserRepository)
	AuthController := controller.NewAuthController(&UserUseCase)

	server.POST("/register", AuthController.Register)
	server.POST("/login", AuthController.Login)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protected := server.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/products", ProductController.GetProducts)

		protected.POST("/product", ProductController.CreateProduct)

		protected.GET("/product/:productId", ProductController.GetProductById)

		protected.PUT("/product/:productId", ProductController.UpdateProduct)

		protected.DELETE("/product/:productId", ProductController.DeleteProduct)
	}

	server.Run(":8000")
}
