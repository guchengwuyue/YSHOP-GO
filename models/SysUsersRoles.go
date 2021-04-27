package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type SysUsersRoles struct {
	Id int64
	UserId *SysUser `orm:"column(user_id);rel(fk)"`
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(SysUsersRoles))
}

func AddUserRole(m *SysUsersRoles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}


