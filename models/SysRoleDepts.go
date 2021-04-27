package models

import "github.com/beego/beego/v2/client/orm"

type SysRolesDepts struct {
	Id int64
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
	DeptId *SysDept `orm:"column(dept_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(SysRolesDepts))
}
