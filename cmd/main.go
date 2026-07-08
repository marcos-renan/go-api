package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	productController := controller.NewProductController()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/products", productController.GetProducts)

	router.Run(":8000") //escuta por padrão na porta 8080
}
