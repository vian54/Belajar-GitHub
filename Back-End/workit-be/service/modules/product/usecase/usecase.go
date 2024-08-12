package usecase

import (
    "github.com/gin-gonic/gin"
    "github.com/ricnah/workit-be/types/models"
    "github.com/ricnah/workit-be/service/extensions/terror"
)

type ProductUsecase interface {
    CreateProduct(ctx *gin.Context, product models.Product) (models.Product, terror.ErrInterface)
    GetProducts(ctx *gin.Context) ([]models.Product, terror.ErrInterface)
    GetProductByID(ctx *gin.Context, id int64) (models.Product, terror.ErrInterface)
    UpdateProduct(ctx *gin.Context, product models.Product) (models.Product, terror.ErrInterface)
    DeleteProduct(ctx *gin.Context, id int64) (terror.ErrInterface)
}
