package admin

import (
	"encoding/json"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
)

//  ConfigController operations for Config
type ConfigController struct {
	BaseController
}

var configDb = new(models.Config)

// Save Post ...
// @Title 批量保存
// @Description  批量保持 Config
// @Param body	body []models.Config true "Config数据"
// @Success 0 {int} models.Config
// @Failure 500 body is empty
// @router /save [post]
func (c *ConfigController) Save() {
	var entity []models.Config
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &entity)
	utils.Nil(err, "接受参数转换错误")
	err = configDb.UpdateAll(entity)
	if err != nil {
		panic(err)
	}
	c.Success("成功")
}

// Find ...
// @Title 获取单个
// @Description 根据用户id查询单个用户信息
// @Param	id		path 	string	true		"Config id"
// @Success 200 {object} models.Config
// @Failure 403 :id is empty
// @router /find/:id [get]
func (c *ConfigController) Find() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Config)
	one, err := entity.GetById(id)
	if err != nil {
		panic(err)
	}
	c.Success(one)
}

// All @Title 获取全部配置
// @Description 获取所有数据
// @Success 200 {object} []models.Config
// @router /list [get]
func (c *ConfigController) All() {
	c.Success(configDb.All())
}
