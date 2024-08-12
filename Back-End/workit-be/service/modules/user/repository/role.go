package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/models"
	"gorm.io/gorm"
)

func (r UserRepository) RoleGetByID(ctx *gin.Context, id int64) (role models.Role, terr terror.ErrInterface) {
	err := r.db.First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}
