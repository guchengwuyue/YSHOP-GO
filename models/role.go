package models

import (
	"github.com/beego/beego/v2/client/orm"
	"yixiang.co/yshop/dto"
	"yixiang.co/yshop/untils"
)

type Role struct {
	Id     int64 `json:"id"`
	Name  string `json:"name" valid:"Required;"`
	Remark string `json:"remark"`
	DataScope string `json:"dataScope"`
	Level int32 `json:"level"`
	Permission string `json:"permission"`
	Users    []*User    `orm:"reverse(many)"`
	Menus []*Menu `json:"menus" orm:"rel(m2m);rel_through(yixiang.co/yshop/models.RolesMenus)"`
	Depts []*Dept `orm:"rel(m2m);rel_through(yixiang.co/yshop/models.RolesDepts)"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Role))
}

func GetOneRole(id int64) Role  {
	o := orm.NewOrm()
	role := Role{Id: id}
	o.Read(&role)
	_, _ = o.LoadRelated(&role, "Menus")
	return role
}

// get all
func GetAllRole(base dto.BasePage,query ...interface{}) (int,[]Role)  {
	var (
		tableName = "role"
		lists []Role
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

func AddRole(m *Role) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByRole(m *Role) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByRole(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE role SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}
