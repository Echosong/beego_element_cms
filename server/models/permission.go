package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Permission struct {
	SysBase
	Name        string `orm:"description(权限名)" json:"name"`
	Description string `orm:"description(描述);size(6000)" json:"description" `
	Url         string `orm:"description(访问路径)" json:"url"`
	Perms       string `orm:"description(权限码)" json:"perms"`
	ParentId    int    `orm:"description(上级id)" json:"parent_id"`
	Type        int8   `orm:"description(类型)" json:"type"`
	Sort        int    `orm:"description(排序)" json:"sort"`
	Icon        string `orm:"description(图标)" json:"icon"`
	Show        bool   `json:"show" orm:"description(显示)"`
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Permission) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// ListByRole  根据 role id 获取权限信息
func (m *Permission) ListByRole(roleId int) (permissions []*Permission) {
	o := orm.NewOrm()
	var rolePermissions []*RolePermission
	o.QueryTable(new(RolePermission)).Filter("role_id", roleId).All(&rolePermissions)
	var ids []int
	for _, permission := range rolePermissions {
		ids = append(ids, permission.PermissionId)
	}
	if len(ids) == 0 {
		return permissions
	}
	o.QueryTable(new(Permission)).Filter("id__in", ids).All(&permissions)
	return permissions
}

// ListByUser 根据用户id 获取当前的权限
func (m *Permission) ListByUser(userId int64) (permissions []*Permission) {
	o := orm.NewOrm()
	user := User{}
	user.Id = userId
	err := o.Read(&user)
	if err != nil {
		panic("用户不存在")
	}
	if user.RoleId == 1 {
		o.QueryTable(new(Permission)).OrderBy("sort").All(&permissions)
	} else {
		var rolePermissions []*RolePermission
		o.QueryTable(new(RolePermission)).Filter("role_id", user.RoleId).All(&rolePermissions)
		var ids []int
		for _, permission := range rolePermissions {
			ids = append(ids, permission.PermissionId)
		}
		if len(ids) == 0 {
			panic("权限不存在")
		}
		o.QueryTable(new(Permission)).Filter("id__in", ids).All(&permissions)
	}
	return permissions
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Permission) GetById(id int64) (v *Permission, err error) {
	o := orm.NewOrm()
	v = &Permission{}
	v.Id = id
	if err = o.QueryTable(new(Permission)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (u *Permission) PageList(query dtos.PermissionQuery) ([]*Permission, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
	qs.Filter("DeleteAt", 0)
	count, _ := qs.Count()
	var list []*Permission
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Permission) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Permission) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
