package repository

import (
    "github.com/ricnah/workit-be/types/models"
    "github.com/ricnah/workit-be/service/extensions/terror"
)

type ProductRepository interface {
    CreateProduct(product *models.Product) terror.ErrInterface
    GetProducts() ([]models.Product, terror.ErrInterface)
    GetProductByID(id int64) (models.Product, terror.ErrInterface)
    UpdateProduct(product *models.Product) terror.ErrInterface
    DeleteProduct(id int64) terror.ErrInterface
}
