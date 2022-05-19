package filters

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"kaiyu-web/models"
	"runtime"
)

func init() {
	web.BConfig.RecoverFunc = RecoverPanic
}

func RecoverPanic(ctx *context.Context, config *web.Config) {
	if err := recover(); err != nil {
		var stack []string
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			//logs.Critical(fmt.Sprintf("%s:%d", file, line))
			stack = append(stack, fmt.Sprintln(fmt.Sprintf("%s:%d", file, line)))
		}
		//自定义的错误，需要写入日志
		logs.Error(err, stack)
		//显示错误
		data := models.ResponseResult{
			Code:    500,
			Message: fmt.Sprintf("%v", err),
			Data:    "",
		}
		_ = ctx.Output.JSON(data, true, true)
	}
}
