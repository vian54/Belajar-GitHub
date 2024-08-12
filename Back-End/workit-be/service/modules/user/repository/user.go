package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/constants"
	"github.com/ricnah/workit-be/types/models"
	"gorm.io/gorm"
)

func (r UserRepository) UserGetByEmail(ctx *gin.Context, email string) (user models.User, terr terror.ErrInterface) {
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r UserRepository) UserGetByPhone(ctx *gin.Context, phone string) (user models.User, terr terror.ErrInterface) {
	err := r.db.First(&user, "phone = ?", phone).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r UserRepository) UserGetByID(ctx *gin.Context, id int64) (user models.User, terr terror.ErrInterface) {
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r UserRepository) UserCreate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface) {
	err := r.db.Create(user).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r UserRepository) UserSearch(ctx *gin.Context, user models.User, searchPayload models.DbSearchObject) (users []models.User, totalData int64, terr terror.ErrInterface) {
	queryDB := r.db.Session(&gorm.Session{})

	// filter
	{
		if user.Email != "" {
			queryDB = queryDB.Where("email like ?", user.Email)
		}
		if user.Name != "" {
			queryDB = queryDB.Where("name like ?", user.Name)
		}
	}

	if searchPayload.Mode == constants.DB_MODE_DATA || searchPayload.Mode == constants.DB_MODE_PAGE {
		offset := (searchPayload.Page - 1) * searchPayload.Limit
		for _, v := range searchPayload.Order {
			queryDB = queryDB.Order(v)
		}
		err := queryDB.Limit(int(searchPayload.Limit)).Offset(int(offset)).Find(&users).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	if searchPayload.Mode == constants.DB_MODE_COUNT || searchPayload.Mode == constants.DB_MODE_PAGE {
		err := queryDB.Model(&models.User{}).Limit(-1).Offset(-1).Count(&totalData).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	return
}

func (r UserRepository) UserUpdate(ctx *gin.Context, user *models.User) (terr terror.ErrInterface) {
	err := r.db.Save(user).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

// custom repos

func (r UserRepository) UserGetByRoleAndFilterName(ctx *gin.Context, roles []string, name string) (users []models.User, terr terror.ErrInterface) {
	err := r.db.Raw(`
		select u.* from users as u inner join (select * from roles where deleted_at is null) as r on r.id = u.role_id
		where r.name in ? and u.name like ? and u.deleted_at is null
	`, roles, name).Scan(&users).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}
