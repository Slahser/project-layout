package main

import (
	"context"

	"github.com/Slahser/coup-de-grace/internal/app/helper"
	"github.com/Slahser/coup-de-grace/internal/pkg/etcd_wrapper"
	"go.uber.org/zap"
)

func main() {

	//全局logger配置
	logger := helper.InitLogger()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	//tt_client_go.Tcg1()

	print(etcd_wrapper.GetEtcdKv().Get(context.TODO(), "/tt"))
}
