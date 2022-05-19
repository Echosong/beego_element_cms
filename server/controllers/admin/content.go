package admin

import (
	"encoding/json"
	"kaiyu-web/dtos"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
	"time"
)

//  ContentController operations for Content
type ContentController struct {
	BaseController
}

// Post ...
// @Title Post
// @Description  新增或者修改 Content
// @Param body	body models.User true "Content数据"
// @Success 0 {int} models.Content
// @Failure 500 body is empty
// @router /save [post]
func (c *ContentController) Save() {
	var entity models.Content
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
// @Param	code		path 	string	true		"Content id"
// @Success 200 {object} models.Content
// @Failure 403 :code is empty
// @router /:code [get]
func (c *ContentController) Find() {
	code := c.GetString(":code")
	utils.IsTrue(code != "", "参数错误")
	entity := new(models.Content)
	one, err := entity.GetByCode(code)
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
func (c *ContentController) ListPage() {
	queryDto := dtos.ContentQuery{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &queryDto)
	if err != nil {
		panic(err)
	}
	entity := &models.Content{}
	var (
		list  []*models.Content
		total int64
	)
	list, total = entity.PageList(queryDto)
	c.Page(list, utils.Pager{Page: queryDto.Page, PageSize: queryDto.PageSize, TotalElements: total})
}

// Delete ...
// @Title Delete
// @Description delete the Content
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ContentController) Delete() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Content)
	_, err := entity.DeleteById(id)
	if err != nil {
		panic(err)
	}
}
