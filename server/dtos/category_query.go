package dtos

import "kaiyu-web/utils"

type CategoryQuery struct {
	utils.Pager
	Type     int `json:"type"`
	ParentId int `json:"parent_id"`
}
