package logging

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context) LoggerInterface {
	logObj := &Logger{
		LogChannel: CHANNEL_RUNTIME,
	}
	if ctx != nil {
		logObj.Ctx = ctx

	}
	return logObj
}

func (l *Logger) Channel(ch string) *Logger {
	if isChannelValidated(ch) {
		l.LogChannel = ch
	}
	return l
}

func (l *Logger) Info(msg interface{}) {
	const level = "info"

	l.WriteLog(level, msg)
}

func (l *Logger) Debug(msg interface{}) {
	const level = "debug"

	l.WriteLog(level, msg)
}

func (l *Logger) Warning(msg interface{}) {
	const level = "warning"

	l.WriteLog(level, msg)
}

func (l *Logger) Error(msg interface{}) {
	const level = "error"

	l.WriteLog(level, msg)
}

func (l *Logger) Fatal(msg interface{}) {
	const level = "fatal"

	l.WriteLog(level, msg)
}

func (l *Logger) Panic(msg interface{}) {
	const level = "panic"

	l.WriteLog(level, msg)
}

func (l *Logger) WriteLog(level string, msg interface{}) {
	fun, file := getCallerFunction()
	go func() {
		logHead := getLogHeader(level)

		var msgString = fmt.Sprintf("%v", msg)
		data, ok := msg.(map[string]interface{})
		if ok {
			for k, v := range data {
				data[k] = v
			}
			if l.LogChannel == CHANNEL_RUNTIME {
				data["function"] = fun
				data["file"] = file

				if l.Ctx != nil {
					reqID := l.Ctx.Value("requestID")
					if reqID != nil {
						data["request_id"] = reqID
					}
				}
			}

			msgByte, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Cannot marshal data")
				return
			}

			msgString = string(msgByte)
		}
		msgString = logHead + msgString
		fmt.Println(msgString)

		logStringDataToFile(l.LogChannel, msgString)
	}()
}
