package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

// Article 资质 案例 主营业务
type Article struct {
	SysBase
	Title      string `orm:"description(标题)" json:"title"`
	Info       string `orm:"description(简介);size(1000)" json:"info"`
	Content    string `orm:"description(内容);type(text)" json:"content"`
	State      int8   `orm:"description(状态1 正常 0 禁用)" json:"state"`
	CategoryId int    `orm:"description(类别)" json:"category_id"`
	Sort       int32  `json:"sort" orm:"description(排序)" json:"sort"`
	Img        string `json:"img" orm:"description(图片)" json:"img"`
	Memo       string `json:"memo" orm:"description(扩展);size(1000)" json:"memo"`
	User       string `json:"user" orm:"description(作者)" json:"user"`
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Article) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Article) GetById(id int64) (v *Article, err error) {
	o := orm.NewOrm()
	v = &Article{}
	v.Id = id
	if err = o.QueryTable(new(Article)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// PageList 查询分页信息
func (u *Article) PageList(query dtos.ArticleQuery) ([]*Article, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	qs = qs.Filter("DeleteAt", 0)

	if query.CategoryId != 0 {
		c := new(Category)
		c.Id = int64(query.CategoryId)
		err := o.Read(c)
		if err != nil {
			return nil, 0
		}
		var ids []int
		ids = append(ids, query.CategoryId)
		if c.ParentId == 0 {
			var cates []*Category
			o.QueryTable(new(Category)).Filter("parent_id", query.CategoryId).All(&cates)
			for _, cate := range cates {
				ids = append(ids, int(cate.Id))
			}
		}
		qs = qs.Filter("category_id__in", ids)
	}
	if len(query.Title) > 1 {
		qs = qs.Filter("title__icontains", query.Title)
	}

	count, _ := qs.Count()
	var list []*Article
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("sort", "-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Article) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Article) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
