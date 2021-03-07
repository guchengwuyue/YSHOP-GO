package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"yixiang.co/yshop/dto"
)

type DictDetail struct {
	Id     int64 `json:"id"`
	Label string `json:"label" valid:"Required;"`
	Value string `json:"value" valid:"Required;"`
	Sort int `json:"sort"`
	DictId int64 `json:"dictId"`
	DictName string `json:"dictName"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(DictDetail))
}

// get all
func GetAllDictDetail(base dto.BasePage,query ...interface{}) (int,[]DictDetail)  {
	var (
		tableName = "dict_detail"
		lists []DictDetail
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

func AddDictDetail(m *DictDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateByDictDetail(m *DictDetail) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(m)
	return
}

func DelByDictDetail(id int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE dict_detail SET is_del = ? WHERE id = ?", 1, id).Exec()
	return
}


