package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"OnlineMeeting/internal/cmd"
	_ "OnlineMeeting/internal/file/logic"
	_ "OnlineMeeting/internal/meeting/logic"
	_ "OnlineMeeting/internal/system/logic"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
