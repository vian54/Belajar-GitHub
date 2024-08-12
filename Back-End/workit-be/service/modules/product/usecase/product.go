package usecase

import (
    "github.com/gin-gonic/gin"
    "github.com/ricnah/workit-be/service/extensions/terror"
    "github.com/ricnah/workit-be/service/modules/product/repository"
    "github.com/ricnah/workit-be/types/models"
)

type productUsecase struct {
    productRepo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
    return &productUsecase{productRepo: repo}
}

func (u *productUsecase) CreateProduct(ctx *gin.Context, product models.Product) (models.Product, terror.ErrInterface) {
    terr := u.productRepo.CreateProduct(&product)
    if terr != nil {
        return models.Product{}, terr
    }
    return product, nil
}

func (u *productUsecase) GetProducts(ctx *gin.Context) ([]models.Product, terror.ErrInterface) {
    products, terr := u.productRepo.GetProducts()
    if terr != nil {
        return nil, terr
    }
    return products, nil
}

func (u *productUsecase) GetProductByID(ctx *gin.Context, id int64) (models.Product, terror.ErrInterface) {
    product, terr := u.productRepo.GetProductByID(id)
    if terr != nil {
        return models.Product{}, terr
    }
    return product, nil
}

func (u *productUsecase) UpdateProduct(ctx *gin.Context, product models.Product) (models.Product, terror.ErrInterface) {
    terr := u.productRepo.UpdateProduct(&product)
    if terr != nil {
        return models.Product{}, terr
    }
    return product, nil
}

func (u *productUsecase) DeleteProduct(ctx *gin.Context, id int64) terror.ErrInterface {
    terr := u.productRepo.DeleteProduct(id)
    if terr != nil {
        return terr
    }
    return nil
}
