package models

import (
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
	util "kaiyu-web/utils"
	"strings"
)

type User struct {
	SysBase
	Username string `orm:"description(账号)" valid:"Required" json:"username"`
	Password string `orm:"description(密码)"  json:"_"`
	Sex      int8   `orm:"description(性别)" json:"sex"`
	Name     string `orm:"description(姓名)" json:"name"`
	State    int8   `orm:"description(状态 0,1)" json:"state"`
	Email    string `orm:"description(邮箱)" json:"email"`
	RegIp    string `orm:"description(注册ip)" json:"reg_ip"`
	LoginIp  string `orm:"description(登陆ip)" json:"login_ip"`
	RoleId   int    `json:"role_id" orm:"description(角色)"`
	RoleName string `json:"role_name" orm:"description(部门)"`
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *User) Add() (int64, error) {
	o := orm.NewOrm()
	srcCode := md5.Sum([]byte(m.Password))
	m.Password = fmt.Sprintf("%x", srcCode)
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *User) GetById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{}
	v.Id = id
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (u *User) PageList(query dtos.UserQuery) ([]*User, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	qs.Filter("DeleteAt", 0)
	count, _ := qs.Count()
	var list []*User
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *User) UpdateById() (int64, error) {
	o := orm.NewOrm()
	one, _ := m.GetById(m.Id)
	if m.Password == "" {
		m.Password = one.Password
	} else {
		srcCode := md5.Sum([]byte(m.Password))
		m.Password = fmt.Sprintf("%x", srcCode)
	}
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *User) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}

// Login 登录处理
func (m *User) Login(login dtos.UserLogin) (err error, user *User) {
	o := orm.NewOrm()
	u := User{}
	u.Username = login.Username
	err = o.Read(&u, "username")
	if err != nil {
		panic("账号或者密码不存在")
	}
	if util.Md5(login.Password) != strings.Trim(u.Password, " ") {
		panic("账号或者密码不存在")
	}
	u.LoginIp = login.LoginIp
	if _, err := o.Update(&u); err != nil {
		panic("登录异常")
	}
	return err, &u
}

// SetPassword 修改密码
func (m *User) SetPassword(dto dtos.PasswordDto) (err error) {
	user, err := m.GetById(int64(dto.UserId))
	if user == nil {
		panic("用户不存在")
	}

	if util.Md5(strings.Trim(dto.Password, " ")) != strings.Trim(user.Password, " ") {
		panic("原始密码错误")
	}
	user.Password = util.Md5(strings.Trim(dto.NewPassword, " "))
	o := orm.NewOrm()
	_, err = o.Update(user)
	if err != nil {
		return err
	}
	return nil

}
