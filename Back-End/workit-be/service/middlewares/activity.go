package middlewares

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func ActivityLogger() gin.HandlerFunc {
	return func(ctxOri *gin.Context) {
		// skip for now
		// {
		// 	ctx := ctxOri.Copy()
		// 	ctxRequest := ctx.Request
		// 	beforeMemory := getCurrentMemoryUsage()
		// 	requestTime := time.Now()
		// 	requestID := "req_" + helper.CreateRandomString(16)

		// 	defer func() {
		// 		if r := recover(); r != nil {
		// 			panicInfo := fmt.Sprint(r) + "\n" + getStackString()
		// 			logging.Log(ctx).Panic(panicInfo)
		// 		}

		// 		afterMemory := getCurrentMemoryUsage()
		// 		memoryUsageDiff := afterMemory - beforeMemory

		// 		activityData := map[string]interface{}{
		// 			"app_host":     ctx.ClientIP(),
		// 			"client_ip":    ctxRequest.RemoteAddr,
		// 			"path":         ctxRequest.URL.Path,
		// 			"requestID":    requestID,
		// 			"agent":        ctxRequest.UserAgent(),
		// 			"responseTime": int64(time.Now().Sub(requestTime).Milliseconds()),
		// 			"httpCode":     ctxOri.Writer.Status(),
		// 			"memoryUsage":  memoryUsageDiff,
		// 			"requestAt":    requestTime.Format("2006-01-02T15:04:05.000Z07:00"),
		// 		}

		// 		logging.Log(ctx).Channel("activity").Info(activityData)
		// 	}()
		// }

		ctxOri.Next()
	}
}

func getCurrentMemoryUsage() uint64 {
	var memStats runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&memStats)
	return memStats.Alloc
}

func getStackString() (res string) {
	stackSize := 64
	stack := make([]uintptr, stackSize)
	length := runtime.Callers(3, stack) // Skip the first 3 frames

	for i := 0; i < length; i++ {
		funcPtr := runtime.FuncForPC(stack[i])
		file, line := funcPtr.FileLine(stack[i])
		res += fmt.Sprintf("%s:%d %s\n", file, line, funcPtr.Name())
	}
	return
}
