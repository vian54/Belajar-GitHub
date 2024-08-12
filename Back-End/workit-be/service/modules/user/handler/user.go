package handler

import (
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/models"
)

func (h UserHandler) UserGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	car, terr := h.userUsecase.UserGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, car)
}

func (h UserHandler) UserGetByEmail(ctx *gin.Context) {
	var (
		emailReq models.EmailRequest
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&emailReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	user, terr := h.userUsecase.UserGetByEmail(ctx, emailReq.Email)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, user)
}

func (h UserHandler) UserCreate(ctx *gin.Context) {
	var (
		user models.User
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	user, terr = h.userUsecase.UserCreate(ctx, user)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, user)
}

func (h UserHandler) UserSearch(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.userUsecase.UserSearch(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
}

func (h UserHandler) UserUpdate(ctx *gin.Context) {
	var (
		user models.User
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	user, terr = h.userUsecase.UserUpdate(ctx, user)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, user)
}

func (h UserHandler) UserGetAllUser(ctx *gin.Context) {
	var (
		user models.UserRole
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.userUsecase.UserGetAllUser(ctx, user.Name)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
}
