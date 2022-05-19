package models

type UserLog struct {
	SysBase
	UserId   string
	UserName string
	Agent    string `orm:"description(客户端信息)"`
	Ip       string `orm:"description(客户端ip)"`
	Url      string `orm:"description(请求url)"`
	Type     int    `orm:"description(业务类型)"`
	Param    string `orm:"description(业务信息);type(text)"`
	Title    string `orm:"description(操作说明)"`
}
