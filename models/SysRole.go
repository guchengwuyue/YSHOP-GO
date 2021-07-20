package models

import (
	"github.com/beego/beego/v2/client/orm"
	"yixiang.co/yshop/common/untils"
	"yixiang.co/yshop/models/dto"
)

type SysRole struct {
	Id     int64        `json:"id"`
	Name  string        `json:"name" valid:"Required;"`
	Remark string       `json:"remark"`
	DataScope string    `json:"dataScope"`
	Level int32         `json:"level"`
	Permission string   `json:"permission"`
	Users    []*SysUser `orm:"reverse(many)"`
	Menus []*SysMenu    `json:"menus" orm:"rel(m2m);rel_through(yixiang.co/yshop/models.SysRolesMenus)"`
	Depts []*SysDept    `orm:"rel(m2m);rel_through(yixiang.co/yshop/models.SysRolesDepts)"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysRole))
}

func GetOneRole(id int64) SysRole {
	o := orm.NewOrm()
	role := SysRole{Id: id}
	o.Read(&role)
	_, _ = o.LoadRelated(&role, "Menus")
	return role
}

// get all
func GetAllRole(base dto.BasePage,query ...interface{}) (int,[]SysRole)  {
	var (
		tableName = "sys_role"
		lists []SysRole
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and name= '" + base.Blurry + "'"
	}


	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&lists)

	o := orm.NewOrm()
	for k, _ := range lists {
		_, _ = o.LoadRelated(&lists[k], "Menus")
	}


	return total,lists
}

func AddRole(m *SysRole) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByRole(m *SysRole) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByRole(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_role SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}
