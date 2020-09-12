package model

import (
	"strconv"
)

type PageOptions struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func ParsePageOptions(limitStr string, offsetStr string, defaultLimit int, defaultOffset int) PageOptions {
	pageOptions := PageOptions{}
	if offsetStr == "" {
		pageOptions.Offset = defaultOffset
	} else {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			panic(err)
		}
		pageOptions.Offset = offset
	}
	if limitStr == "" {
		pageOptions.Limit = defaultLimit
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			panic(err)
		}
		pageOptions.Limit = limit
	}
	return pageOptions
}
