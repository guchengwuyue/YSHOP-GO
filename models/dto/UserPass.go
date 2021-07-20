package dto

type UserPass struct {
	OldPass string `json:"oldPass" valid:"Required;"`
	NewPass string `json:"newPass" valid:"Required;"`
}
