package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"系统管理" method:"post" summary:"用户登录"`
	UserID   string `v:"required#用户ID不能为空"`
	UserName string `v:"required#用户名不能为空"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"`
}
