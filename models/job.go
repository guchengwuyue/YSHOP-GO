package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"yixiang.co/yshop/dto"
	"yixiang.co/yshop/untils"
)

type Job struct {
	Id     int64 `json:"id"`
	Name string `json:"name" valid:"Required;"`
	Enabled int8 `json:"enabled"`
	Sort int8 `json:"sort"`
	//DeptId int64 `json:"deptId"`
	Dept *Dept `json:"dept" orm:"column(dept_id);bigint;rel(one)""`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Job))
}


// get all
func GetAllJob(base dto.BasePage,query ...interface{}) (int,[]Job)  {
	var (
		tableName = "job"
		lists []Job
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and name= '" + base.Blurry + "'"
	}
	if len(query) > 0 {
		enabled := query[0].(int64)
		if enabled >= 0 {
			condition += " and enabled=" + strconv.FormatInt(enabled,10)
		}
	}

	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&lists)

	o := orm.NewOrm()
	for k, _ := range lists {
		_, _ = o.LoadRelated(&lists[k], "Dept")
	}


	return total,lists
}

func AddJob(m *Job) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByJob(m *Job) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByJob(ids []int64) (err error) {
	str := untils.ReturnQ(len(ids))
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE job SET is_del = ? WHERE id in("+str+")", 1, ids).Exec()
	return
}
