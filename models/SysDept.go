package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/common/untils"
)

type SysDept struct {
	Id     int64       `json:"id"`
	Name string        `json:"name" valid:"Required;"`
	Pid int64          `json:"pid"`
	Enabled int8       `json:"enabled" valid:"Required;"`
	Children []SysDept `orm:"-" json:"children"`
	Label string       `orm:"-" json:"label"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysDept))
}

func GetAllDepts(name string, enabled int8) []SysDept {
	var depts []SysDept
	o := orm.NewOrm()
	qs := o.QueryTable("sys_dept").Filter("is_del",0)
	if name != "" {
		qs = qs.Filter("name",name)
	}
	if enabled > -1 {
		qs = qs.Filter("enabled",enabled)
	}
	qs.All(&depts)
	return RecursionDeptList(depts,0)
}

//递归函数
func RecursionDeptList(data []SysDept, pid int64) []SysDept {
	var listTree = make([]SysDept,0)
	for _, value := range data {
		value.Label = value.Name
		if value.Pid == pid {
			value.Children = RecursionDeptList(data, value.Id)
			listTree = append(listTree, value)
		}
	}
	return listTree
}

func AddDept(m *SysDept) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByDept(m *SysDept) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByDept(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	logs.Info(str)
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_dept SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}
