package dto

type UserPost struct {
	Id     int64 `json:"id"`
	NickName string `json:"nickName" valid:"Required;"`
	Phone string `json:"phone" valid:"Required;"`
	Sex string `json:"sex"`
}