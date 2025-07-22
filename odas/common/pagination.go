package common

import (
	"net/url"
	"strconv"
)

// Pagination 分页信息
type Pagination struct {
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
	Total    int64 `json:"total,omitempty"`
	Pages    int64 `json:"pages,omitempty"`
}

func (p Pagination) Params(prefix string) url.Values {
	params := url.Values{}
	if p.Page > 0 {
		params.Add("page", strconv.FormatInt(p.Page, 10))
	}
	if p.PageSize > 0 {
		params.Add("pageSize", strconv.FormatInt(p.PageSize, 10))
	}
	return params
}

// Page 分页信息
type Page struct {
	Page     int64 `json:"page,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
}

func (p Page) Params(prefix string) url.Values {
	params := url.Values{}
	if p.Page > 0 {
		params.Add("page", strconv.FormatInt(p.Page, 10))
	}
	if p.PageSize > 0 {
		params.Add("pageSize", strconv.FormatInt(p.PageSize, 10))
	}
	return params
}
