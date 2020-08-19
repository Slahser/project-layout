package app

import (
	"time"

	"github.com/Slahser/coup-de-grace/pkg/ttserver/path"

	"github.com/Slahser/coup-de-grace/internal/app/helper"
	endless "github.com/fvbock/endless"
	gin "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	listeningPort = ":8080"
	runningMode   = gin.DebugMode
)

func Run() error {

	gin.SetMode(runningMode)
	router := gin.New()
	//基础中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//入参校验
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zap.S().Info("入参校验 test log here")
		_ = v.RegisterValidation("ttValidation", helper.TtValidation)
	}

	//路由组织
	funcRouter := router.Group("/funcs")
	{
		funcRouter.GET("/:funcName", path.TtWeb)
	}

	//server微调
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	return endless.ListenAndServe(listeningPort, router)
}
