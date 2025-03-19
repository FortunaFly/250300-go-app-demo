package main

import (
	"fmt"
	"os"

	"github.com/lwmacct/250300-go-app-demo/app"
	"github.com/lwmacct/250300-go-app-demo/app/start"
	"github.com/lwmacct/250300-go-app-demo/app/test"

	"github.com/lwmacct/250300-go-app-demo/app/version"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
)

var mc *mflag.Ts

func main() {
	mc = mflag.New(nil)

	{
		// 命令行参数
		mc.AddCobra(version.Cmd().Cobra())
		mc.AddCobra(start.Cmd().Cobra())

		// 开发环境中的测试命令
		if os.Getenv("ACF_SHOW_TEST") == "1" {
			mc.AddCobra(test.Cmd().Cobra())
		}
	}

	{
		mlog.SetNew(
			mlog.WithFile(app.Flag.Log.File),
			mlog.WithLevel(app.Flag.Log.Level),
			mlog.WithCallerClip(version.Workspace),
		)
	}

	if err := mc.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
