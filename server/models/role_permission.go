package models

import "github.com/beego/beego/v2/client/orm"

type RolePermission struct {
	SysBase
	RoleId       int `orm:"description(角色id)"`
	PermissionId int `orm:"description(权限id)"`
}

// Bind 绑定信息
func (r RolePermission) Bind(permissionIds []int, roleId int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(RolePermission)).Filter("role_id", roleId).Delete()
	if err != nil {
		return err
	}
	var rs []*RolePermission
	for _, id := range permissionIds {
		p := new(RolePermission)
		p.PermissionId = id
		p.RoleId = roleId
		rs = append(rs, p)
	}
	_, err = o.InsertMulti(len(rs), rs)
	return err
}
