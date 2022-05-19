package admin

import (
	"encoding/json"
	"kaiyu-web/dtos"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
	"time"
)

//  CategoryController operations for Category
type CategoryController struct {
	BaseController
}

// Post ...
// @Title Post
// @Description  新增或者修改 Category
// @Param body	body models.User true "Category数据"
// @Success 0 {int} models.Category
// @Failure 500 body is empty
// @router /save [post]
func (c *CategoryController) Save() {
	var entity models.Category
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
	c.Success("成功")
}

// Find ...
// @Title 获取单个
// @Description 根据用户id查询单个用户信息
// @Param	id		path 	string	true		"Category id"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /find/:id [get]
func (c *CategoryController) Find() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Category)
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
func (c *CategoryController) ListPage() {
	queryDto := dtos.CategoryQuery{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &queryDto)
	if err != nil {
		panic(err)
	}
	entity := &models.Category{}
	var (
		list  []*models.Category
		total int64
	)
	list, total = entity.PageList(queryDto)
	c.Page(list, utils.Pager{Page: queryDto.Page, PageSize: queryDto.PageSize, TotalElements: total})
}

// Delete ...
// @Title Delete
// @Description delete the Category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CategoryController) Delete() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Category)
	_, err := entity.DeleteById(id)
	if err != nil {
		panic(err)
	}
	c.Success("成功")
}
