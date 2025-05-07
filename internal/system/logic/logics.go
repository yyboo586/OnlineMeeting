package logic

import (
	"OnlineMeeting/internal/system/logic/middleware"
	"OnlineMeeting/internal/system/logic/token"
	"OnlineMeeting/internal/system/service"
)

func init() {
	service.RegisterMiddleware(middleware.New())
	service.RegisterGToken(token.New())
}
