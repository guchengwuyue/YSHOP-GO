package models

import "github.com/beego/beego/v2/client/orm"

type Picture struct {
	Id     int64
	DeleteUrl string
	Filename string
	Height string
	Size string
	Url string
	Username string
	Width string
	Md5code string
	BaseModel
}

func init() {
	orm.RegisterModel(new(Picture))
}
