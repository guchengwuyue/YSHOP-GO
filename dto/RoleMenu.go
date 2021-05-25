package dto

type RoleMenu struct {
	Id int64     `json:"id"`
	Menus []Menu `json:"menus" valid:"Required;"`
}

type Menu struct {
	Id int64 `json:"id"`
}
