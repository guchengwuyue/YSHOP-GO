package dto

type AuthUser struct {
	Code string `json:"code"`
	Password string `json:"password"`
	Username string `json:"username"`
	Id          string `json:"uuid"`
}
