package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func UserCreateRepository(db *gorm.DB) IRepository {
	userRepository := UserRepository{
		db: db,
	}
	return userRepository
}

type IRepository interface {
	RoleGetByID(ctx *gin.Context, id int64) (role models.Role, terr terror.ErrInterface)
	UserGetByEmail(ctx *gin.Context, email string) (user models.User, terr terror.ErrInterface)
	UserGetByPhone(ctx *gin.Context, phone string) (user models.User, terr terror.ErrInterface)
	UserGetByID(ctx *gin.Context, id int64) (user models.User, terr terror.ErrInterface)
	UserCreate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface)
	UserSearch(ctx *gin.Context, user models.User, searchPayload models.DbSearchObject) (users []models.User, totalData int64, terr terror.ErrInterface)
	UserUpdate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface)
	UserGetByRoleAndFilterName(ctx *gin.Context, roles []string, name string) (users []models.User, terr terror.ErrInterface)
}
