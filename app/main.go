package main

import (
	"net/http"

	"github.com/aldysp34/simple-crud/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Success")
	})

	productRepo := controllers.New()
	r.POST("/product", productRepo.CreateProduct)
	r.GET("/products", productRepo.GetProducts)
	r.GET("/product/:id", productRepo.GetProduct)
	r.PUT("/product/:id", productRepo.UpdateProduct)
	r.DELETE("/product/:id", productRepo.DeleteProduct)

	return r
}
