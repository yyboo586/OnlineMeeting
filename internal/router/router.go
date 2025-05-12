package router

import (
	"context"

	fileRouter "OnlineMeeting/internal/file/router"
	meetingRouter "OnlineMeeting/internal/meeting/router"
	systemRouter "OnlineMeeting/internal/system/router"

	systemService "OnlineMeeting/internal/system/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Router struct {
}

var R = new(Router)

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1/online_meeting", func(group *ghttp.RouterGroup) {
		// 封装GoFrame格式的返回数据
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Middleware(systemService.Middleware().MiddlewareCORS)
		group.Middleware(systemService.Middleware().Ctx)
		// 绑定system模块路由
		systemRouter.R.BindController(ctx, group)
		// 绑定meeting模块路由
		meetingRouter.R.BindController(ctx, group)
		// 绑定file模块路由
		fileRouter.R.BindController(ctx, group)
	})
}
