package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Product struct {
	Title      string `json:"name" orm:"description(标题)"`
	Content    string `json:"content" orm:"description(内容);type(text)"`
	Info       string `json:"info" orm:"description(简介)"`
	Img        string `json:"img" orm:"description(图片)"`
	CategoryId int    `json:"category_id" orm:"description(分类 1产品 2 合作伙伴 3 行业客户)"`
	SysBase
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Product) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// Save 保持新
func (m *Product) Save(one *Product) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
	//传图片形式
	if one.CategoryId != 24 {
		count, _ := qs.Filter("category_id", one.CategoryId).Count()
		if count == 0 {
			_, err = o.Insert(one)
		} else {
			_, err = o.Update(one)
		}
	} else {
		if one.Id == 0 {
			_, err = o.Insert(one)
		} else {
			_, err = o.Update(one)
		}
	}
	return err
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Product) GetById(id int64) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{}
	v.Id = id
	if err = o.QueryTable(new(Product)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// ListByCategory 根据分类查询信息
func (p Product) ListByCategory(categoryIds []int) (list []*Product, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Product)).Filter("category_id__in", categoryIds).All(&list)
	return list, err
}

// PageList 查询分页信息
func (u *Product) PageList(query dtos.ProductQuery) ([]*Product, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
	if len(query.Title) > 1 {
		qs = qs.Filter("title__icontains", query.Title)
	}
	count, _ := qs.Count()
	var list []*Product
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Product) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Product) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
