package dtos

import "kaiyu-web/utils"

type UserQuery struct {
	utils.Pager
	Username string `json:"username" description:"账号"` //账号
	State    int    `json:"state" description:"状态"`    //状态
}
