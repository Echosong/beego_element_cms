package dtos

import "kaiyu-web/utils"

type AdQuery struct {
	utils.Pager
	CategoryId int `json:"category_id"`
	Type       int `json:"type"`
}
