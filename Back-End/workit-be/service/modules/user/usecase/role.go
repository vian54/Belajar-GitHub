package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/models"
)

func (u UserUsecase) RoleGetByID(ctx *gin.Context, id int64) (role models.Role, terr terror.ErrInterface) {
	return u.userRepo.RoleGetByID(ctx, id)
}
