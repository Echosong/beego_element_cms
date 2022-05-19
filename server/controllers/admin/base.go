package admin

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/server/web"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"math"
	"reflect"
	"strings"
)

const (
	responseCode = 200
)

type BaseController struct {
	web.Controller
}

// Prepare 初始化链接数据
func (c BaseController) Prepare() {
	_, actionName := c.GetControllerAndAction()
	if strings.ToLower(actionName) != "login" {
		if c.GetSession("user") == nil {
			c.Alert("未登录", 410)
		}
	}
	logs.Info("开始执行.....登录成功....")
}

func (c BaseController) Finish() {
	logs.Info("完成执行....")
}

func (c *BaseController) Success(o interface{}) {
	c.Alert(o, responseCode)
}

func (c *BaseController) Alert(o interface{}, code int64) {
	c.Data["json"] = &models.ResponseResult{Data: o, Code: code, Success: code == responseCode}
	c.ServeJSON()
}

func (c *BaseController) Page(list interface{}, pager utils.Pager) {
	pager.TotalPages = int(math.Ceil(float64(pager.TotalElements) / float64(pager.PageSize)))
	rv := reflect.ValueOf(list)
	if rv.IsNil() {
		list = new([0]int)
	}
	c.Data["json"] = &models.ResponseResult{
		Data:     list,
		Code:     responseCode,
		Success:  true,
		PageData: pager,
	}
	c.ServeJSON()
}

//获取用户IP地址
func (p *BaseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
