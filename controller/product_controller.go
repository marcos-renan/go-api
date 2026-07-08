package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct{}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "HyperX Cloud III",
			Price: 449.99,
		},
		{
			ID:    2,
			Name:  "Fortrek BlackFire",
			Price: 149.99,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
