package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Role struct {
	SysBase
	Name        string `orm:"description(角色名称)" json:"name"`
	Description string `orm:"description(描述)" json:"description"`
	Code        string `json:"code" orm:"description(编码)"`
	ParentId    int    `orm:"description(上级id)" json:"parent_id"`
	Sort        int    `json:"sort" orm:"description(排序)"`
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Role) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Role) GetById(id int64) (v *Role, err error) {

	o := orm.NewOrm()
	v = &Role{}
	v.Id = id
	if err = o.QueryTable(new(Role)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (u *Role) PageList(query dtos.RoleQuery) ([]*Role, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
	qs.Filter("DeleteAt", 0)
	count, _ := qs.Count()
	var list []*Role
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Role) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Role) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
