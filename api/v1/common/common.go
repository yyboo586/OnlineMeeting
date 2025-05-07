package common

import "github.com/gogf/gf/v2/frame/g"

type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}

type PageReq struct {
	PageNum  int `dc:"当前页码"`
	PageSize int `dc:"每页数"`
	// OrderBy  string `dc:"排序方式"`
}

type PageRes struct {
	CurrentPage int
	Total       int
}
