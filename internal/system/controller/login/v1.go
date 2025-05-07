package login

import (
	"OnlineMeeting/api/v1/system"
	"OnlineMeeting/internal/system/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type loginControllerV1 struct{}

func NewV1() service.ILoginV1 {
	return &loginControllerV1{}
}

func (c *loginControllerV1) Login(ctx context.Context, req *system.LoginReq) (res *system.LoginRes, err error) {
	key := "sdssssssssssssssssssssssssssssssssssssssssssstom"
	userInfo := g.Map{
		"UserID":   req.UserID,
		"UserName": req.UserName,
	}
	data, err := service.GfToken().GenerateToken(ctx, key, userInfo)
	if err != nil {
		return
	}

	res = new(system.LoginRes)
	res.Token = data
	return
}
