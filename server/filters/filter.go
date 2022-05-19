package filters

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"kaiyu-web/models"
)

func init() {
	beego.InsertFilterChain("/*", func(next beego.FilterFunc) beego.FilterFunc {
		return func(ctx *context.Context) {
			// do something

			// don't forget this
			next(ctx)
			// 如果是操作类型的就没必要输出，没有输出的时候就认为是成功的
			result := models.ResponseResult{
				Code:    200,
				Success: true,
				Message: "成功2",
			}
			ctx.Output.JSON(result, true, true)

		}
	})
}
