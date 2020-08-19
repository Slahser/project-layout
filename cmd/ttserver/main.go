package ttserver

import (
	"os"

	ttserverapp "github.com/Slahser/coup-de-grace/cmd/ttserver/app"

	"github.com/Slahser/coup-de-grace/internal/app/helper"
	"go.uber.org/zap"
)

func main() {

	//全局logger配置
	logger := helper.InitLogger()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	//web server启动
	if err := ttserverapp.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
