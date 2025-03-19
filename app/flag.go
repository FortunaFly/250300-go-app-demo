package app

import "github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"

type TsFlag struct {
	Log   mlog.Opts
	Start struct{} `group:"start" note:"默认配置"`

	App struct {
		ListenAddr string `group:"app" note:"监听地址" default:"0.0.0.0:8888"`
	}
}
