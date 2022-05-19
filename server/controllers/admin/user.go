package admin

import (
	"encoding/base64"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"kaiyu-web/dtos"
	"kaiyu-web/models"
	"kaiyu-web/utils"
	"strconv"
	"time"
)

// UserController  operations for User
type UserController struct {
	BaseController
}

// Save Post ...
// @Title Post
// @Description  新增或者修改 User
// @Param body	body models.User true "User数据"
// @Success 0 {int} models.User
// @Failure 500 body is empty
// @router /save [post]
func (c *UserController) Save() {
	var entity models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &entity)
	roleDb := new(models.Role)
	role, err := roleDb.GetById(int64(entity.RoleId))
	if err != nil {
		panic("请选择部门角色")
	}
	entity.RoleName = role.Name
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
// @Param	id		path 	string	true		"User id"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /find/:id [get]
func (c *UserController) Find() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.User)
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
func (c *UserController) ListPage() {
	queryDto := dtos.UserQuery{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &queryDto)
	if err != nil {
		panic(err)
	}
	entity := &models.User{}
	var (
		list  []*models.User
		total int64
	)
	list, total = entity.PageList(queryDto)
	c.Page(list, utils.Pager{Page: queryDto.Page, PageSize: queryDto.PageSize, TotalElements: total})
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	id, _ := strconv.ParseInt(c.GetString(":id"), 0, 64)
	utils.IsTrue(id != 0, "参数错误")
	entity := new(models.User)
	_, err := entity.DeleteById(id)
	if err != nil {
		panic(err)
	}
}

// GetCurrent @Title 获取当前用户信息
// @router /getCurrent [get]
func (m *UserController) GetCurrent() {
	m.Success(m.GetSession("user"))
}

// Login @Title 登录
// @router /login [post]
func (c *UserController) Login() {
	loginDto := dtos.UserLogin{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginDto)
	if err != nil {
		panic(err)
	}
	data, _ := base64.StdEncoding.DecodeString(loginDto.Password)
	loginDto.Password = string(data)
	loginDto.LoginIp = c.getClientIp()
	userDb := new(models.User)
	var user *models.User
	err, user = userDb.Login(loginDto)
	if err != nil {
		panic("登录异常")
	}
	c.SetSession("user", user)
	c.Success(utils.UniqueId())
}

// Logout @Tile 登录退出
// @router /logout [get]
func (c *UserController) Logout() {
	c.DestroySession()
	c.Success("退出登录")
}

// SetPassword 修改密码
// @router /setPassword [post]
func (c *UserController) SetPassword() {
	passwordDto := dtos.PasswordDto{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &passwordDto)
	if err != nil {
		panic(err)
	}
	if passwordDto.RePassword != passwordDto.NewPassword {
		panic("确认密码不一样")
	}
	obj := c.GetSession("user")
	logs.Info(obj)
	user := obj.(*models.User)
	if user.Id == 0 {
		panic("未登录")
	}
	passwordDto.UserId = int(user.Id)
	userDb := new(models.User)
	err = userDb.SetPassword(passwordDto)
	if err != nil {
		panic(err)
	}
	c.Success("修改密码成功")
}
