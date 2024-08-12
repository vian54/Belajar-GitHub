package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/constants"
	"github.com/ricnah/workit-be/types/models"
	"golang.org/x/crypto/bcrypt"
)

func (u UserUsecase) AuthGetFromContext(ctx *gin.Context) (res models.UserRole, terr terror.ErrInterface) {
	session := ctx.Value("user_id")
	userID, ok := session.(int64)
	if !ok {
		terr = terror.New(errors.New("Cannot get user from session"))
		return
	}

	if userID <= 0 {
		terr = terror.ErrInvalidRule("User in Session Not Found")
		return
	}

	userRes, terr := u.UserGetByID(ctx, userID)
	if terr != nil {
		terr = terror.ErrInvalidRule("User in Session Not Found")
		return
	}

	r, terr := u.RoleGetByID(ctx, userRes.RoleID)
	if terr != nil {
		terr = terror.ErrInvalidRule("Role not found for user")
		return
	}

	utstruct.InjectStructValue(userRes, &res)
	res.RoleName = r.Name

	return
}

func (u UserUsecase) AuthLogin(ctx *gin.Context, username string, password string) (authResp models.AuthResponse, terr terror.ErrInterface) {
	var (
		errx terror.ErrInterface
		user models.User
	)

	user, errx = u.UserGetByEmail(ctx, username)
	if errx != nil {
		if errx.GetType() == terror.ERROR_TYPE_DATA_NOT_FOUND {
			user, errx = u.UserGetByPhone(ctx, username)
			if errx != nil {
				if errx.GetType() == terror.ERROR_TYPE_DATA_NOT_FOUND {
					terr = terror.ErrInvalidRule("User with the email / phone was not found")
					return
				} else {
					terr = errx
					return
				}
			}
		} else {
			terr = errx
			return
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		terr = terror.ErrInvalidRule("User Password is not match")
		return
	}

	var (
		tokenString string
		expires     time.Time
	)
	// token generation
	{
		expires = time.Now().Add(time.Hour * 3)

		// Create the JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
			Issuer:    utstring.GetEnv(constants.ENV_APP_NAME),
			Subject:   fmt.Sprintf("%d", user.ID),
		})

		// Sign the token with a secret key
		tokenString, err = token.SignedString([]byte(utstring.GetEnv(constants.ENV_APP_SECRET, "")))
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	r, terr := u.RoleGetByID(ctx, user.RoleID)
	if terr != nil {
		terr = terror.ErrInvalidRule("Role not found for user")
		return
	}

	authResp = models.AuthResponse{
		User:      user,
		Token:     tokenString,
		Role:      r,
		ExpiredAt: expires.Format(time.RFC3339),
	}

	return
}
