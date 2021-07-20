package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"yixiang.co/yshop/models/dto"
)

type SysDictDetail struct {
	Id     int64 `json:"id"`
	Label string `json:"label" valid:"Required;"`
	Value string `json:"value" valid:"Required;"`
	Sort int `json:"sort"`
	DictId int64 `json:"dictId"`
	DictName string `json:"dictName"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(SysDictDetail))
}

// get all
func GetAllDictDetail(base dto.BasePage,query ...interface{}) (int,[]SysDictDetail)  {
	var (
		tableName = "sys_dict_detail"
		lists []SysDictDetail
		condition = ""
	)
	if base.Blurry != "" {
		condition = " and label= '" + base.Blurry + "'"
	}
	if len(query) > 0 {
		dictId := query[0].(int64)
		if dictId > 0 {
			condition += " and dict_id=" + strconv.FormatInt(dictId,10)
		}
		if len(query) > 1 {
			condition += " and dict_name= '" + query[1].(string) + "'"
		}
	}

	total, _, rs := GetPagesInfo(tableName, base.Page, base.Size, condition)
	rs.QueryRows(&lists)


	return total,lists
}

func AddDictDetail(m *SysDictDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByDictDetail(m *SysDictDetail) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByDictDetail(id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE sys_dict_detail SET is_del = ? WHERE id = ?", 1, id).Exec()
	return
}


