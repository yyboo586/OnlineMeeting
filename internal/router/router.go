package router

import (
	"context"

	meetingRouter "OnlineMeeting/internal/meeting/router"
	systemRouter "OnlineMeeting/internal/system/router"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Router struct {
}

var R = new(Router)

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		// 封装GoFrame格式的返回数据
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 绑定system模块路由
		systemRouter.R.BindController(ctx, group)
		// 绑定meeting模块路由
		meetingRouter.R.BindController(ctx, group)
	})
}
