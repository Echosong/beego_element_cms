package admin

import (
	"encoding/json"
	"kaiyu-web/dtos"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
	"time"
)

//  PermissionController operations for Permission
type PermissionController struct {
	BaseController
}

var permissionDb = new(models.Permission)

// ListByUser
// @Title 根据用户获取当前的权限
// @Description 获取当前用户的权限
// @Success {object} [] models.Permission
// @router /listByUser/:id [get]
func (p *PermissionController) ListByUser() {
	id, _ := strconv.Atoi(p.GetString(":id"))
	if id == 0 {
		id = 1
	}
	p.Success(permissionDb.ListByUser(int64(id)))
}

// ListByRole @Title 根据角色查看权限
// @router /listByRole/:id [get]
func (p *PermissionController) ListByRole() {
	id, _ := strconv.Atoi(p.GetString(":id"))
	if id == 0 {
		id = 1
	}
	p.Success(permissionDb.ListByRole(id))
}

// UpdateRolePermissions 绑定信息
// @router /updateRolePermissions/:id [post]
func (p *PermissionController) UpdateRolePermissions() {
	id, _ := strconv.Atoi(p.GetString(":id"))
	if id == 0 {
		id = 1
	}
	var ids []int
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &ids)
	if err != nil {
		panic("参数错误")
	}
	m := new(models.RolePermission)
	err = m.Bind(ids, id)
	if err != nil {
		panic(err)
	}
	p.Success("")

}

// Post ...
// @Title Post
// @Description  新增或者修改 Permission
// @Param body	body models.User true "Permission数据"
// @Success 0 {int} models.Permission
// @Failure 500 body is empty
// @router /save [post]
func (c *PermissionController) Save() {
	var entity models.Permission
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
}

// Find ...
// @Title 获取单个
// @Description 根据用户id查询单个用户信息
// @Param	id		path 	string	true		"Permission id"
// @Success 200 {object} models.Permission
// @Failure 403 :id is empty
// @router /find/:id [get]
func (c *PermissionController) Find() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Permission)
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
func (c *PermissionController) ListPage() {
	queryDto := dtos.PermissionQuery{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &queryDto)
	if err != nil {
		panic(err)
	}
	entity := &models.Permission{}
	var (
		list  []*models.Permission
		total int64
	)
	list, total = entity.PageList(queryDto)
	c.Page(list, utils.Pager{Page: queryDto.Page, PageSize: queryDto.PageSize, TotalElements: total})
}

// Delete ...
// @Title Delete
// @Description delete the Permission
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PermissionController) Delete() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.Permission)
	_, err := entity.DeleteById(id)
	if err != nil {
		panic(err)
	}
}
