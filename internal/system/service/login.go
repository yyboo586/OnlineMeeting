package service

import (
	"OnlineMeeting/api/v1/system"
	"context"
)

type ILoginV1 interface {
	Login(ctx context.Context, req *system.LoginReq) (res *system.LoginRes, err error)
}
