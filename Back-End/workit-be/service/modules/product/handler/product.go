package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/models"
)

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	var terr terror.ErrInterface

	if err := c.ShouldBindJSON(&product); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(c, terr)
		return
	}

	createdProduct, terr := h.productUsecase.CreateProduct(c, product)
	if terr != nil {
		ResponseJson(c, terr)
		return
	}

	ResponseJson(c, createdProduct)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, terr := h.productUsecase.GetProducts(c)
	if terr != nil {
		ResponseJson(c, terr)
		return
	}

	ResponseJson(c, products)
}
