package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

type BaseModel struct {
	UpdateTime time.Time `orm:"auto_now_add;type(datetime)" json:"updateTime"`
	CreateTime time.Time `orm:"auto_now;type(datetime)" json:"createTime"`
	IsDel int8 `json:"isDel"`
}

/**
* 分页函数，适用任何表
* 返回 总记录条数,总页数,以及当前请求的数据RawSeter,调用中需要"rs.QueryRows(&tblog)"就行了  --tblog是一个Tb_log对象
* 参数：表名，当前页数，页面大小，条件（查询条件,格式为 " and name='zhifeiya' and age=12 "）
*/
func GetPagesInfo(tableName string, currentpage int, pagesize int, conditions string) (int, int, orm.RawSeter) {
	if currentpage <= 1 {
		currentpage = 1
	}
	if pagesize == 0 {
		pagesize = 20
	}
	var rs orm.RawSeter
	o := orm.NewOrm()
	var totalItem, totalpages int = 0, 0                                                          //总条数,总页数
	o.Raw("SELECT count(*) FROM " + tableName + "  where is_del=0 " + conditions).QueryRow(&totalItem) //获取总条数
	if totalItem <= pagesize {
		totalpages = 1
	} else if totalItem > pagesize {
		temp := totalItem / pagesize
		if (totalItem % pagesize) != 0 {
			temp = temp + 1
		}
		totalpages = temp
	}
	rs = o.Raw("select *  from  " + tableName + "  where is_del=0 " + conditions +  " order by id desc " +
		" LIMIT " + strconv.Itoa((currentpage-1)*pagesize) + "," + strconv.Itoa(pagesize))
	return totalItem, totalpages, rs
}


