package dto

type UserQuery struct {
	BasePage
	Sort string
	Blurry string
	Enabled bool
}