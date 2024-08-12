package models

import (
	"time"

	"gorm.io/gorm"
)

type DefaultResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Response struct {
	ResponseCode  string      `json:"responseCode"`
	ResponseDesc  string      `json:"responseDesc"`
	ResponseData  interface{} `json:"responseData"`
	ResponseTrace string      `json:"responseTrace"`
}

type DbStandard struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
