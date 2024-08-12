package product

import (
    "github.com/gin-gonic/gin"
    "github.com/ricnah/workit-be/service/modules/product/handler"

)

func InitRoutes(v1 *gin.RouterGroup, productHandler *handler.ProductHandler) {
    productRoute := v1.Group("/products")
    {
        productRoute.POST("/create", productHandler.CreateProduct)
        productRoute.GET("/getlist", productHandler.GetProducts)
    }
}
