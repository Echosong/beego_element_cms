package models

import "time"

// SysBase 数据表基类
type SysBase struct {
	Id         int64     `json:"id"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"`
	UpdateTime time.Time `orm:"auto_now_add;type(datetime)" json:"update_time"`
	CreatorId  int       `orm:"default(0);description(添加)" json:"creator_id"`
	UpdaterId  int       `orm:"default(0);description(更新人)" json:"updater_id"`
	DeleteAt   int8      `orm:"default(0);description(软删除状态)" json:"delete_at"`
}

// ResponseResult 返回模型
type ResponseResult struct {
	Code     int64       `json:"code"`
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	PageData interface{} `json:"page_data"`
}
