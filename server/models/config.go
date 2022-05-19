package models

import (
	"github.com/beego/beego/v2/client/orm"
	"kaiyu-web/dtos"
)

type Config struct {
	Key         string `json:"key" orm:"description(键)"`
	Name        string `json:"name" orm:"description(名称)"`
	Value       string `json:"value" orm:"description(值)"`
	Group       string `json:"group" orm:"description(分组)"`
	Type        int    `json:"type" orm:"description(类型)"`
	Description string `json:"description" orm:"description(描述)"`
	SysBase
}

// Add AddUser insert a new User into database and returns
// last inserted Id on success.
func (m *Config) Add() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetById GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func (m *Config) GetById(id int64) (v *Config, err error) {
	o := orm.NewOrm()
	v = &Config{}
	v.Id = id
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateAll 更新全部
func (m *Config) UpdateAll(list []Config) (err error) {
	o := orm.NewOrm()
	listOld := m.All()
	for _, config := range list {
		for _, c := range listOld {
			//更新
			if c.Group == config.Group && c.Key == config.Key && c.Value != config.Value {
				_, err = o.Update(&config)
				break
			}
		}
	}
	return err
}

// GetConfigs 获取配置文件信息
func (m *Config) GetConfigs(group string) map[string]interface{} {
	list := m.All()
	var mp = make(map[string]interface{})
	for _, config := range list {
		if group == config.Group {
			mp[config.Key] = config.Value
		}
	}
	return mp
}

func (m *Config) All() (list []*Config) {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Config)).All(&list)
	if err != nil {
		return nil
	}
	return list
}

func (m *Config) PageList(query dtos.ConfigQuery) ([]*Config, int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Config))
	qs.Filter("DeleteAt", 0)
	count, _ := qs.Count()
	var list []*Config
	if count > 0 {
		offset := (query.Page - 1) * query.PageSize
		qs.OrderBy("-id").Limit(query.PageSize, offset).All(&list)
	}
	return list, count
}

// UpdateById UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func (m *Config) UpdateById() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

// DeleteById DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func (m *Config) DeleteById(id int64) (int64, error) {
	o := orm.NewOrm()
	m.Id = id
	return o.Delete(m)
}
