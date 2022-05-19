package admin

import (
	"encoding/json"
	"kaiyu-web/dtos"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
	"time"
)

//  AdController operations for Ad
type AdController struct {
	BaseController
}

// Post ...
// @Title Post
// @Description  新增或者修改 Ad
// @Param body	body models.User true "Ad数据"
// @Success 0 {int} models.Ad
// @Failure 500 body is empty
// @router /save [post]
func (c *AdController) Save() {
	var entity models.Ad
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &entity)
	entity.UpdateTime = time.Now()
	utils.Nil(err, "接受参数转换错误")
	if entity.Id > 0 {
		_, err = entity.UpdateById()
	} else {
		_, err = entity.Add()
	}
	if err != nil {
		panic(err)
	}
	c.Success("")
}

// Find ...
// @Title 获取单个
// @Description 根据用户id查询单个用户信息
// @Param	id		path 	string	true		"Ad id"
// @Success 200 {object} models.Ad
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AdController) Find() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Ad)
	one, err := entity.GetById(id)
	if err != nil {
		panic(err)
	}
	c.Success(one)
}

// ListPage ...
// @Title 分页查询
// @Description 分页查询
// @Param	userQuery	body 	dtos.UserQuery	true		"查询列表"
// @Success 200 {object} []models.User
// @Failure 500 empty
// @router /listPage [put]
func (c *AdController) ListPage() {
	queryDto := dtos.AdQuery{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &queryDto)
	if err != nil {
		panic(err)
	}
	entity := &models.Ad{}
	var (
		list  []*models.Ad
		total int64
	)
	list, total = entity.PageList(queryDto)
	c.Page(list, utils.Pager{Page: queryDto.Page, PageSize: queryDto.PageSize, TotalElements: total})
}

// Delete ...
// @Title Delete
// @Description delete the Ad
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AdController) Delete() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Ad)
	_, err := entity.DeleteById(id)
	if err != nil {
		panic(err)
	}
	c.Success("")
}
