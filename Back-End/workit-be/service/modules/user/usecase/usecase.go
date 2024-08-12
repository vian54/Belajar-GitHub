package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/service/modules/user/repository"
	"github.com/ricnah/workit-be/types/models"
)

type UserUsecase struct {
	userRepo repository.IRepository
}

func UserCreateUsecase(userRepo repository.IRepository) IUsecase {
	userUsecase := UserUsecase{
		userRepo: userRepo,
	}
	return userUsecase
}

type IUsecase interface {
	AuthGetFromContext(ctx *gin.Context) (res models.UserRole, terr terror.ErrInterface)
	AuthLogin(ctx *gin.Context, email string, password string) (authResp models.AuthResponse, terr terror.ErrInterface)
	RoleGetByID(ctx *gin.Context, id int64) (role models.Role, terr terror.ErrInterface)
	UserGetByEmail(ctx *gin.Context, email string) (user models.User, terr terror.ErrInterface)
	UserGetByPhone(ctx *gin.Context, phone string) (user models.User, terr terror.ErrInterface)
	UserGetByID(ctx *gin.Context, id int64) (user models.User, terr terror.ErrInterface)
	UserUpdate(ctx *gin.Context, user models.User) (userRes models.User, terr terror.ErrInterface)
	UserCreate(ctx *gin.Context, user models.User) (userRes models.User, terr terror.ErrInterface)
	UserSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface)
	UserGetAllUser(ctx *gin.Context, name string) (users []models.User, terr terror.ErrInterface)
}
