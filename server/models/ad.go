package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Ad struct {
	SysBase
	Title      string `json:"title" orm:"description(标题)"`
	Img        string `json:"img" orm:"description(路径)"`
	Link       string `json:"link" orm:"description(链接)"`
	State      int8   `json:"state" orm:"description(状态)" `
	Sort       int32  `json:"sort" orm:"description(排序)"`
	CategoryId int    `json:"category_id" orm:"description(分类)" type:"input"`
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Ad) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Ad) GetById(id int64) (v *Ad, err error) {
	o := orm.NewOrm()
	v = &Ad{}
	v.Id = id
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (u *Ad) PageList(query dtos.AdQuery) ([]*Ad, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ad))
	qs = qs.Filter("DeleteAt", 0)
	if 0 != query.CategoryId {
		qs = qs.Filter("category_id", query.CategoryId)
	}
	if 0 != query.Type {
		c := new(Category)
		var ids []int
		var categories []*Category
		o.QueryTable(c).Filter("type", query.Type).All(&categories, "id")
		for _, category := range categories {
			ids = append(ids, int(category.Id))
		}
		qs.Filter("category_id__in", ids)
	}
	count, _ := qs.Count()
	var list []*Ad
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Ad) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Ad) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
