package model_view

import "goblog/model"

// BlogView parses a create or update blog request
// it consists of a blog and a list of tags
type BlogView struct {
	Blog model.Blog `json:"blog"`
	Type model.Type `json:"type"`
	Tags []model.Tag `json:"tags"`
}