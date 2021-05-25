package models

import (
	"github.com/beego/beego/v2/client/orm"
	"yixiang.co/yshop/dto"
)

type SysRolesMenus struct {
	Id int64
	MenuId *SysMenu `orm:"column(menu_id);rel(fk)"`
	RoleId *SysRole `orm:"column(role_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(SysRolesMenus))
}

func BatchRoleMenuAdd(menu dto.RoleMenu)  {
	o := orm.NewOrm()
	o.Raw("delete from sys_roles_menus WHERE role_id = ?",  menu.Id).Exec()

	var roleMenus []SysRolesMenus
	for _, val := range menu.Menus {
		var menus = SysMenu{Id: val.Id}
		var roles = SysRole{Id: menu.Id}
		roleMenus = append(roleMenus, SysRolesMenus{MenuId: &menus,RoleId: &roles})
	}

	o.InsertMulti(100,roleMenus)
}
