package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/helper"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/constants"
	"github.com/ricnah/workit-be/types/models"
	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) UserGetByEmail(ctx *gin.Context, email string) (user models.User, terr terror.ErrInterface) {
	return u.userRepo.UserGetByEmail(ctx, email)
}

func (u UserUsecase) UserGetByPhone(ctx *gin.Context, phone string) (user models.User, terr terror.ErrInterface) {
	return u.userRepo.UserGetByPhone(ctx, phone)
}

func (u UserUsecase) UserGetByID(ctx *gin.Context, id int64) (user models.User, terr terror.ErrInterface) {
	return u.userRepo.UserGetByID(ctx, id)
}

func (u UserUsecase) UserUpdate(ctx *gin.Context, user models.User) (userRes models.User, terr terror.ErrInterface) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			terr = terror.New(err)
			return
		}
		user.Password = string(hashedPassword)
	}
	user.Email = ""

	terr = u.userRepo.UserUpdate(ctx, &user)
	if terr != nil {
		return
	}

	userRes = user
	return
}

func (u UserUsecase) UserCreate(ctx *gin.Context, user models.User) (userRes models.User, terr terror.ErrInterface) {
	_, totalData, terr := u.userRepo.UserSearch(ctx, models.User{}, models.DbSearchObject{
		Mode: constants.DB_MODE_COUNT,
	})
	if terr != nil {
		return
	}

	if totalData > 0 {
		var operator models.UserRole
		operator, terr = u.AuthGetFromContext(ctx)
		if terr != nil {
			return
		}

		if operator.RoleName != string(constants.ROLES_ADMIN) {
			terr = terror.ErrInvalidRule("not allowed roled to create user")
			return
		}
	}

	// check existing user
	{
		_, terr = u.UserGetByEmail(ctx, user.Email)
		if terr == nil {
			terr = terror.ErrInvalidRule("User with the email has been exist")
			return
		}
		if terr.GetType() != terror.ERROR_TYPE_DATA_NOT_FOUND {
			return
		}
		terr = nil
	}

	// start create user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		terr = terror.New(err)
		return
	}
	user.Password = string(hashedPassword)

	terr = u.userRepo.UserCreate(ctx, &user)
	if terr != nil {
		return
	}

	userRes = user
	return
}

func (u UserUsecase) UserSearch(ctx *gin.Context, filter models.DbSearchObject) (res models.DbSearchObject, terr terror.ErrInterface) {
	filter.Mode = constants.DB_MODE_PAGE

	user := models.User{}
	err := helper.MapAnyToStruct(filter.PayloadData, &user)
	if err != nil {
		terr = terror.New(err)
		return
	}

	err = helper.WrapPercentOnStructString(&user)
	if err != nil {
		terr = terror.New(err)
		return
	}

	usersRes, totalData, terr := u.userRepo.UserSearch(ctx, user, filter)
	if terr != nil {
		return
	}

	filter.ResponseData = usersRes
	filter.TotalData = totalData
	res = filter

	return
}

func (u UserUsecase) UserGetAllUser(ctx *gin.Context, name string) (users []models.User, terr terror.ErrInterface) {
	name = helper.WrapString(name, "%")
	return u.userRepo.UserGetByRoleAndFilterName(ctx, []string{string(constants.ROLES_ADMIN), string(constants.ROLES_USER)}, name)
}
