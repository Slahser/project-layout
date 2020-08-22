package main

import (
	"github.com/Slahser/coup-de-grace/internal/app/helper"
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	//全局logger配置
	logger := helper.InitLogger()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	middlewareList, _ := traefikwrapper.Middlewares().List(metav1.ListOptions{})

	print(middlewareList)
}
