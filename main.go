package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"OnlineMeeting/internal/cmd"
	_ "OnlineMeeting/internal/system/logic/token"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
