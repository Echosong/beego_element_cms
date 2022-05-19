package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Operate struct {
	SysBase
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Operate) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (u *Operate) GetById(id int64) (v *Operate, err error) {
	o := orm.NewOrm()
	v = &Operate{}
	v.Id = id
	if err = o.QueryTable(new(Operate)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (u *Operate) PageList(query dtos.OperateQuery) ([]*Operate, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Operate))
	qs.Filter("DeleteAt", 0)
	count, _ := qs.Count()
	var list []*Operate
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Operate) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Operate) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
