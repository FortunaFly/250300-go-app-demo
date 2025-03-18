package app

import "github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"

type TsFlag struct {
	Log   mlog.Opts
	Start struct{} `group:"start" note:"默认配置"`

	App struct {
		Port int `group:"app" note:"Http 服务端口" default:"8888"`
	}
}
