package dtos

import "kaiyu-web/utils"

type ProductQuery struct {
	utils.Pager
	Title string `json:"title"`
}
