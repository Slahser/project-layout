package goproxy_selfhost

import (
	"github.com/goproxy/goproxy"
	"net/http"
	"os"
)

func SelfHostGoProxy(){
	g := goproxy.New()
	g.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy_selfhost.cn,direct", // 使用 goproxy_selfhost.cn 作为上游代理
		"GOPRIVATE=git.example.com",         // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
	)
	g.ProxiedSUMDBs = []string{"sum.golang.org https://goproxy_selfhost.cn/sumdb/sum.golang.org"} // 代理默认的校验和数据库
	_ = http.ListenAndServe("localhost:10086", g)
}
