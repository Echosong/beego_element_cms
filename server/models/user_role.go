package models

type UserRole struct {
	SysBase
	RoleId int `orm:"description(角色id)"`
	UserId int `orm:"description(用户id)"`
}
