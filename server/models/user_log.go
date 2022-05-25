package models

type UserLog struct {
	SysBase
	UserId   string
	UserName string
	Agent    string `orm:"description(客户端信息)" json:"agent"`
	Ip       string `orm:"description(客户端ip)" json:"ip"`
	Url      string `orm:"description(请求url)" json:"url"`
	Type     int    `orm:"description(业务类型)" json:"type"`
	Param    string `orm:"description(业务信息);type(text)" json:"param" query:"no"`
	Title    string `orm:"description(操作说明)" json:"title" query:"list"`
}
