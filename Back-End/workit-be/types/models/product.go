package models

type Product struct {
	DbStandard
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
}
