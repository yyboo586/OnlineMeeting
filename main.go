package main

import (
	_ "OnlineMeeting/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"OnlineMeeting/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
