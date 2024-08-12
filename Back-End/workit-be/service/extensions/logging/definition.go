package logging

import (
	"fmt"
	nativeLog "log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/gin-gonic/gin"
	"github.com/ricnah/workit-be/types/constants"
)

const (
	CHANNEL_ACTIVITY = "activity"
	CHANNEL_RUNTIME  = "runtime"
	CHANNEL_WORKER   = "worker"
)

var CHANNELS = []string{CHANNEL_ACTIVITY, CHANNEL_RUNTIME, CHANNEL_WORKER}

type Logger struct {
	Ctx        *gin.Context
	LogChannel string
}

type LoggerInterface interface {
	Channel(channel string) *Logger
	Info(data interface{})
	Debug(data interface{})
	Warning(data interface{})
	Error(data interface{})
	Fatal(data interface{})
	Panic(data interface{})
}

func isChannelValidated(chn string) bool {
	for _, v := range CHANNELS {
		if v == chn {
			return true
		}
	}
	return false
}

func getCallerFunction() (function string, file string) {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return fn, filename
}

func getLogHeader(level string) string {
	appName := toCamelCase(utstring.GetEnv(constants.ENV_APP_NAME, ""))
	appEnv := utstring.GetEnv(constants.ENV_APP_ENV, "local")
	logLevel := strings.ToUpper(level)
	timeNow := time.Now().Format("2006-01-02T15:04:05.000Z07:00")

	return fmt.Sprintf("[%s] ::%s.%s.%s:: ", timeNow, appName, appEnv, logLevel)
}

func toCamelCase(input string) (camelCase string) {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '_'
	})

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	camelCase = strings.Join(words, "")

	return
}

// Log To Logs Directory
//
// Log data to storage/logs/filename file
func logStringDataToFile(fileName string, data string) (err error) {
	fileName = strings.TrimSpace(fileName)
	if !strings.HasSuffix(fileName, ".log") {
		fileName += ".log"
	}

	var file *os.File
	logsDir := "./storage/logs"
	_, errStat := os.Stat(logsDir)
	if os.IsNotExist(errStat) {
		err = os.Mkdir(logsDir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	filePath := fmt.Sprintf("./storage/logs/%s", fileName)
	_, errStat = os.Stat(filePath)
	if os.IsNotExist(errStat) {
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		file.Close()
	}

	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	logger := nativeLog.New(file, "", 0)
	logger.Print(data)
	return
}
