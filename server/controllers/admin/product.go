package admin

import (
	"encoding/json"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
	"strings"
	"time"
)

//  ProductController operations for Product
type ProductController struct {
	BaseController
}

// SaveList @Title 参数信息
// @router /save [post]
func (c *ProductController) SaveList() {
	var entitys []*models.Product
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &entitys)
	utils.Nil(err, "接受参数转换错误")
	productDb := new(models.Product)
	for _, entity := range entitys {
		err := productDb.Save(entity)
		if err != nil {
			panic(err)
		}
	}
	c.Success("")
}

// Post ...
// @Title Post
// @Description  新增或者修改 Product
// @Param body	body models.User true "Product数据"
// @Success 0 {int} models.Product
// @Failure 500 body is empty
// @router / [post]
func (c *ProductController) Save() {
	var entity models.Product
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
// @Param	id		path 	string	true		"Product id"
// @Success 200 {object} models.Product
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductController) Find() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Product)
	one, err := entity.GetById(id)
	if err != nil {
		panic(err)
	}
	c.Success(one)
}

// List ...
// @Title 根据分类查询
// @Description 根据分类查询
// @Param	userQuery	body 	string	true		"查询列表"
// @Success 200 {object} []models.User
// @Failure 500 empty
// @router /list/:ids [get]
func (c *ProductController) List() {
	categoryStr := c.GetString(":ids")
	productDb := new(models.Product)
	split := strings.Split(categoryStr, ",")
	var categoryIds []int
	for _, s := range split {
		atom, _ := strconv.Atoi(s)
		categoryIds = append(categoryIds, atom)
	}
	category, err := productDb.ListByCategory(categoryIds)
	if err != nil {
		panic(err)
	}
	c.Success(category)
}

// Delete ...
// @Title Delete
// @Description delete the Product
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductController) Delete() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Product)
	_, err := entity.DeleteById(id)
	if err != nil {
		panic(err)
	}
	c.Success("")
}
