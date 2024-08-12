package repository

import (
    "errors"

    "github.com/ricnah/workit-be/service/extensions/terror"
    "github.com/ricnah/workit-be/types/models"
    "gorm.io/gorm"
)

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *models.Product) (terr terror.ErrInterface) {
    err := r.db.Create(product).Error
    if err != nil {
        terr = terror.New(err)
    }
    return
}

func (r *productRepository) GetProducts() (products []models.Product, terr terror.ErrInterface) {
    err := r.db.Find(&products).Error
    if err != nil {
        terr = terror.New(err)
    }
    return
}

func (r *productRepository) GetProductByID(id int64) (product models.Product, terr terror.ErrInterface) {
    err := r.db.First(&product, id).Error
    if err != nil {
        if errors.Is(gorm.ErrRecordNotFound, err) {
            terr = terror.ErrNotFoundData(err.Error())
            return
        }
        terr = terror.New(err)
    }
    return
}

func (r *productRepository) UpdateProduct(product *models.Product) (terr terror.ErrInterface) {
    err := r.db.Save(product).Error
    if err != nil {
        if errors.Is(gorm.ErrRecordNotFound, err) {
            terr = terror.ErrNotFoundData(err.Error())
            return
        }
        terr = terror.New(err)
    }
    return
}

func (r *productRepository) DeleteProduct(id int64) (terr terror.ErrInterface) {
    err := r.db.Delete(&models.Product{}, id).Error
    if err != nil {
        if errors.Is(gorm.ErrRecordNotFound, err) {
            terr = terror.ErrNotFoundData(err.Error())
            return
        }
        terr = terror.New(err)
    }
    return
}
