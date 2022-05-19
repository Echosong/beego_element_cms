package dtos

import "kaiyu-web/utils"

type ArticleQuery struct {
	utils.Pager
	CategoryId int    `json:"category_id"`
	Title      string `json:"title"`
}
