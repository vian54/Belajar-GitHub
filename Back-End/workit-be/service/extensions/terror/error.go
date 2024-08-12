package terror

import (
	"errors"
	"fmt"
	nativeRuntime "runtime"
	"strings"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/ricnah/workit-be/types/constants"
)

type (
	ErrorModel struct {
		Code    string `json:"responseCode"`
		Type    string `json:"responseDesc"`
		Message string `json:"responseData"`
		Trace   string `json:"responseTrace"`
	}

	errorResponse struct {
		Status int               `json:"status"`
		Data   errorResponseData `json:"data"`
	}

	errorResponseData struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Trace   string `json:"trace"`
	}

	ErrInterface interface {
		GetMessage() string
		GetType() string
		GetNativeError() error
	}
)

func New(err error) ErrInterface {
	var errS ErrorModel
	var traceStack string
	if utstring.GetEnv(constants.ENV_APP_ENV, "") != "production" {
		traceStack = getErrorFlow()
	}

	errS.Code = ERROR_CODE_SYSTEM
	errS.Type = ERROR_TYPE_SYSTEM
	errS.Message = ERROR_DEFAULT_MESSAGE_SYSTEM
	if err != nil {
		errS.Message = err.Error()
	}
	errS.Trace = traceStack
	return &errS
}

func ErrInvalidRule(message ...string) ErrInterface {
	var errS ErrorModel
	var traceStack string
	if utstring.GetEnv(constants.ENV_APP_ENV, "") != "production" {
		traceStack = getErrorFlow()
	}

	errS.Code = ERROR_CODE_INVALID_RULE
	errS.Type = ERROR_TYPE_INVALID_RULE
	errS.Message = ERROR_DEFAULT_MESSAGE_INVALID_RULE
	if len(message) > 0 {
		errS.Message = message[0]
	}
	errS.Trace = traceStack
	return &errS
}

func ErrParameter(message ...string) ErrInterface {
	var errS ErrorModel
	var traceStack string
	if utstring.GetEnv(constants.ENV_APP_ENV, "") != "production" {
		traceStack = getErrorFlow()
	}

	errS.Code = ERROR_CODE_PARAMETER
	errS.Type = ERROR_TYPE_PARAMETER
	errS.Message = ERROR_DEFAULT_MESSAGE_PARAMETER
	if len(message) > 0 {
		errS.Message = message[0]
	}
	errS.Trace = traceStack
	return &errS
}

func ErrNotFoundData(message ...string) ErrInterface {
	var errS ErrorModel
	var traceStack string
	if utstring.GetEnv(constants.ENV_APP_ENV, "") != "production" {
		traceStack = getErrorFlow()
	}

	errS.Code = ERROR_CODE_DATA_NOT_FOUND
	errS.Type = ERROR_TYPE_DATA_NOT_FOUND
	errS.Message = ERROR_DEFAULT_MESSAGE_DATA_NOT_FOUND
	if len(message) > 0 {
		errS.Message = message[0]
	}
	errS.Trace = traceStack
	return &errS
}

func (er *ErrorModel) GetMessage() string {
	return er.Message
}

func getErrorFlow() (rowsJoin string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured", r)
		}
	}()

	pc := make([]uintptr, 10)
	n := nativeRuntime.Callers(3, pc)
	frames := nativeRuntime.CallersFrames(pc[:n])

	rootMandatoryDirs := []string{"/service/"}

	getIfStringContainsMandatoryDir := func(str string) string {
		for _, v := range rootMandatoryDirs {
			if strings.Contains(str, v) {
				return v
			}
		}
		return ""
	}

	for {
		frame, more := frames.Next()

		rootDir := getIfStringContainsMandatoryDir(frame.Function)
		if rootDir == "" {
			break
		}

		var functionName string
		fFunctions := strings.Split(frame.Function, "/")
		functionNames := fFunctions[len(fFunctions)-1]
		functionNameSegments := strings.Split(functionNames, ".")
		functionName = functionNameSegments[len(functionNameSegments)-1]

		mandStrIndex := strings.Index(frame.File, rootDir)
		if mandStrIndex < 0 {
			break
		}
		fileName := frame.File[mandStrIndex:]

		lineNumb := frame.Line

		rowsJoin += fmt.Sprintf("%s(line:%d) on function %s", fileName, lineNumb, functionName)

		if !more {
			break
		}
		rowsJoin += " <- "
	}
	rowsJoin = strings.TrimSuffix(rowsJoin, " <- ")

	return
}

func (i *ErrorModel) GetNativeError() error {
	return errors.New(i.GetMessage())
}

func (i *ErrorModel) GetType() string {
	return i.Type
}
