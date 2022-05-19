package dtos

import "kaiyu-web/utils"

type RoleQuery struct {
	Name string `json:"name"`
	utils.Pager
}
