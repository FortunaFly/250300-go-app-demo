package start

import (
	"fmt"
	"net/http"

	"github.com/lwmacct/250300-go-app-demo/app"

	"github.com/lwmacct/250300-go-mod-mflag/pkg/mflag"
	"github.com/lwmacct/250300-go-mod-mlog/pkg/mlog"
	"github.com/spf13/cobra"
)

func Cmd() *mflag.Ts {
	mc := mflag.New(app.Flag).UsePackageName("")
	mc.AddCmd(func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	}, "run", "", "app", "mlog")
	return mc
}

func run(cmd *cobra.Command, args []string) {
	_ = map[string]any{"cmd": cmd, "args": args}
	mlog.Info(mlog.H{"msg": "app.Flag", "data": app.Flag})

	// 设置 HTTP 处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// 启动 HTTP 服务器
	port := ":34188"
	mlog.Info(mlog.H{"msg": "Starting HTTP server", "port": port})
	if err := http.ListenAndServe(port, nil); err != nil {
		mlog.Error(mlog.H{"msg": "HTTP server error", "error": err})
	}

	mlog.Close()
}
