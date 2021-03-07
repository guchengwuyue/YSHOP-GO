package vo

import "yixiang.co/yshop/models"

type LoginVo struct {
	Token string `json:"token"`
	User *models.User `json:"user"`
}
