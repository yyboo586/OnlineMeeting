package router

import (
	"context"

	controller "OnlineMeeting/internal/meeting/controller/meeting"
	systemService "OnlineMeeting/internal/system/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/meetings", func(group *ghttp.RouterGroup) {
		//登录验证拦截
		systemService.GfToken().Middleware(group)
		group.Bind(
			controller.MeetingController,
		)
	})
}
