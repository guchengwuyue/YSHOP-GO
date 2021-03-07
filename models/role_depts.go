package models

import "github.com/beego/beego/v2/client/orm"

type RolesDepts struct {
	Id int64
	RoleId *Role `orm:"column(role_id);rel(fk)"`
	DeptId *Dept `orm:"column(dept_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(RolesDepts))
}
