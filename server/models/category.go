package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Category struct {
	SysBase
	Name     string `orm:"description(名称)" json:"name"`
	ParentId int    `orm:"description(上级编号)" json:"parent_id"`
	Sort     int    `json:"sort" orm:"decription(排序)"`
	Banner   string `json:"img" orm:"description(banner图片);size(800)"`
	Type     int    `orm:"description(类型 1 友情链接 3 资质)" json:"type"`
	Url      string `json:"url" orm:"description(链接)""`
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Category) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Category) GetById(id int64) (v *Category, err error) {
	o := orm.NewOrm()
	v = &Category{}
	v.Id = id
	if err = o.QueryTable(new(Category)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// PageList 分页获取信息
func (u *Category) PageList(query dtos.CategoryQuery) ([]*Category, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Category))
	qs = qs.Filter("Type", query.Type)
	if 0 != query.ParentId {
		qs = qs.Filter("parent_id", query.ParentId)
	}
	count, _ := qs.Count()
	var list []*Category
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("sort").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Category) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Category) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
