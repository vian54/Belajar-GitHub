package helper

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/types/constants"
	"gorm.io/gorm"
)

func WrapPercentOnStructString(data interface{}) (err error) {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		err = errors.New("Input must be a pointer to a struct")
		return
	}

	structVal := val.Elem()

	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Field(i)

		if field.Kind() == reflect.String {
			currentValue := field.String()

			if currentValue != "" {
				uppercaseValue := WrapString(currentValue, "%")
				field.SetString(uppercaseValue)
			}
		}
	}

	return
}

type TransactionSetFunc = func(ctx *gin.Context) interface{}

func TxCreate(ctx *gin.Context, fn TransactionSetFunc) {
	if TxGet(ctx) == nil {
		txCtx := fn(ctx)
		var ok bool
		var tx *gorm.DB
		tx, ok = txCtx.(*gorm.DB)
		if !ok {
			return
		}
		ctx.Set(constants.TX_CTX_KEY, tx.Begin())
	}
}

func TxGet(ctx *gin.Context) (tx *gorm.DB) {
	txCtx, exist := ctx.Get(constants.TX_CTX_KEY)
	if !exist {
		return
	}
	var ok bool
	tx, ok = txCtx.(*gorm.DB)
	if !ok {
		return nil
	}
	return
}

func TxCommit(ctx *gin.Context) {
	txCtx, exist := ctx.Get(constants.TX_CTX_KEY)
	if !exist {
		return
	}
	tx, ok := txCtx.(*gorm.DB)
	if !ok {
		return
	}
	tx.Commit()
}

func TxRollBack(ctx *gin.Context) {
	txCtx, exist := ctx.Get(constants.TX_CTX_KEY)
	if !exist {
		return
	}
	tx, ok := txCtx.(*gorm.DB)
	if !ok {
		return
	}
	tx.Rollback()
}
