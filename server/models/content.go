package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Content struct {
	Title   string `json:"title" orm:"description(标题)"`
	Img     string `json:"img" orm:"description(banner图片)"`
	TopImg  string `json:"top_img" orm:"description(导航图片)"`
	Content string `json:"content" orm:"description(内容);type(text)"`
	Bs      string `json:"bs" orm:"description(中部业务);type(text)"`
	Video   string `json:"video" orm:"description(视频地址)"`
	Code    string `json:"code" orm:"description(编码)"`
	SysBase
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Content) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Content) GetById(id int64) (v *Content, err error) {
	o := orm.NewOrm()
	v = &Content{}
	v.Id = id
	if err = o.QueryTable(new(Content)).Filter("Id", id).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetByCode 根据code 获取单个信息
func (m *Content) GetByCode(code string) (v *Content, err error) {
	o := orm.NewOrm()
	v = &Content{Code: code}
	err = o.QueryTable(new(Content)).Filter("code", code).One(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (u *Content) PageList(query dtos.ContentQuery) ([]*Content, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Content))
	qs.Filter("DeleteAt", 0)
	count, _ := qs.Count()
	var list []*Content
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Content) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Content) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
