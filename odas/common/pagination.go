package common

import (
	"net/http"
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

func (p Pagination) Body() []byte {
	return nil
}

func (p Pagination) Method() string {
	return http.MethodGet
}

func (p Pagination) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (p Pagination) AuthRequired() bool {
	return true
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

func (p Page) Body() []byte {
	return nil
}

func (p Page) Method() string {
	return http.MethodGet
}

func (p Page) ContentType() string {
	return "application/x-www-form-urlencoded"
}

func (p Page) AuthRequired() bool {
	return true
}
