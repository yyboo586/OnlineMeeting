package login

import (
	"OnlineMeeting/api/v1/system"
	"OnlineMeeting/internal/system/service"
	"OnlineMeeting/library/liberr"
	"context"
)

type loginControllerV1 struct{}

func NewV1() service.ILoginV1 {
	return &loginControllerV1{}
}

func (c *loginControllerV1) Login(ctx context.Context, req *system.LoginReq) (res *system.LoginRes, err error) {
	key := "sdssssssssssssssssssssssssssssssssssssssssssstom"
	data, err := service.GfToken().GenerateToken(ctx, key, req.UserID)
	liberr.ErrIsNil(ctx, err, "生成token失败")

	res = new(system.LoginRes)
	res.Token = data
	return
}
