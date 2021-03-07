package models

import (
	"github.com/beego/beego/v2/client/orm"
	"yixiang.co/yshop/dto"
)

type RolesMenus struct {
	Id int64
	MenuId *Menu `orm:"column(menu_id);rel(fk)"`
	RoleId *Role `orm:"column(role_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(RolesMenus))
}

func BatchRoleMenuAdd(menu dto.RoleMenu)  {
	o := orm.NewOrm()
	o.Raw("delete from roles_menus WHERE role_id = ?",  menu.Id).Exec()

	var roleMenus []RolesMenus
	for _, val := range menu.Menus {
		var menus = Menu{Id: val.Id}
		var roles = Role{Id:menu.Id}
		roleMenus = append(roleMenus,RolesMenus{MenuId: &menus,RoleId: &roles})
	}

	o.InsertMulti(100,roleMenus)
}
