package middlewares

import (
	"fmt"
	"net/http"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/service/extensions/helper"
	"github.com/ricnah/workit-be/service/extensions/terror"
	"github.com/ricnah/workit-be/types/constants"
	"github.com/ricnah/workit-be/types/models"
)

func responseJson(ctx *gin.Context, data interface{}) {
	httpStatusCode := http.StatusOK
	resp, ok := data.(*terror.ErrorModel)
	if ok {
		responseData := models.Response{
			ResponseDesc: resp.Message,
			ResponseCode: resp.Code,
		}

		appEnv := utstring.GetEnv(constants.ENV_APP_ENV, "local")
		if appEnv != "production" {
			responseData.ResponseTrace = resp.Trace
		}

		ctx.JSON(httpStatusCode, responseData)
		return
	}

	responseData := models.Response{
		ResponseDesc: "Success",
		ResponseCode: "00",
	}

	if helper.IsStruct(data) || helper.IsMap(data) {
		responseData.ResponseData = data
	} else {
		responseData.ResponseDesc = fmt.Sprintf("%v", data)
	}

	ctx.JSON(httpStatusCode, responseData)
	return
}
