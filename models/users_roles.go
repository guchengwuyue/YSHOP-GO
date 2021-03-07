package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type UsersRoles struct {
	Id int64
	UserId *User `orm:"column(user_id);rel(fk)"`
	RoleId *Role `orm:"column(role_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(UsersRoles))
}

func AddUserRole(m *UsersRoles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}


