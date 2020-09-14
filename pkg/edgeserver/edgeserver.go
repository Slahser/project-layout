package edgeserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fvbock/endless"
	ginlimits "github.com/gin-contrib/size"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/gorilla/websocket"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"

	ophelper "aliyun.com/opur/opur/pkg/helper"
)

/**
摘抄一段FC的自定义Server.

监听在任何IP（0.0.0.0）的指定端口（端口可以读取环境变量FC_SERVER_PORT，默认为9000）。
HTTP Server需配置connection keep-alive。
请求超时时间设置为15分钟以上。
HTTP Server需要在25秒内启动完毕。
*/

var (
	ListeningPort = ":8080"
	RunningMode   = gin.DebugMode

	ginModeArray = []string{
		gin.DebugMode,
		gin.TestMode,
		gin.ReleaseMode,
	}

	Router *gin.Engine

	UpGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//path /user/:name
	TemplateHTTPGetHandlerFunc gin.HandlerFunc = func(context *gin.Context) {
		name := context.Param("name")
		zap.S().Info("name is : " + name)
		context.String(http.StatusOK, "Hello %s", name)
	}

	//body name=manu&message=this_is_great
	TemplateHTTPPostHandlerFunc gin.HandlerFunc = func(context *gin.Context) {
		//POSIX
		//sh.Exec(context.Param("dir"))
		name := context.PostForm("name")
		message := context.PostForm("message")
		zap.S().Debug("name is : " + name)
		context.JSON(http.StatusOK, name+message)
	}

	//path /ping
	TemplateWSGetHandlerFunc gin.HandlerFunc = func(context *gin.Context) {
		wsConn := GetWSConn(context)
		for {
			mt, message, err := wsConn.ReadMessage()
			if err != nil {
				break
			}
			if string(message) == "ping" {
				message = []byte("pong")
			}
			zap.S().Info("websocket write back pong")
			err = wsConn.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
)

type EdgeServer struct {
	config Config
}

type Config struct {
	MetaConfig
	//LogConfig Log
	ServerConfig
}

type MetaConfig struct {
	Env      string
	Org      string
	Project  string
	FuncName string
}

type ServerConfig struct {
	ListeningPort int
	RunningMode   string
}

func NewForConfigOrDie(config *Config) *EdgeServer {
	server, err := NewForConfig(config)
	if err != nil {
		panic(errors.Errorf("generate edge server error %w", err))
	}
	return server
}

func NewForConfig(config *Config) (*EdgeServer, error) {
	c := *config

	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj 下,端口是否冲突

	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/port
	// v=> config.ServerConfig.ListeningPort

	server := &EdgeServer{c}
	server.initRouteAndMw()
	return server, nil
}

func (server *EdgeServer) initRouteAndMw() {
	//TODO
	zap.ReplaceGlobals(ophelper.InitLogger())
	if server.config.RunningMode != "" && funk.ContainsString(ginModeArray, server.config.RunningMode) {
		gin.SetMode(server.config.RunningMode)
	} else {
		gin.SetMode(RunningMode)
	}
	Router = gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(ginlimits.RequestSizeLimiter(10))
	Router.Use(ginzap.Ginzap(ophelper.Logger, time.RFC3339, true))
	Router.Use(ginzap.RecoveryWithZap(ophelper.Logger, true))
	//Router.NoMethod(NoMethodHandler)
	//Router.NoRoute(NoRouteHandler)
}

func GetWSConn(c *gin.Context) *websocket.Conn {
	wsConn, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(errors.Errorf("get websocket connection error %w", err))
	}
	return wsConn
}

func (server *EdgeServer) Start() {

	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20

	if server.config.ListeningPort != 0 {
		ListeningPort = ":" + strconv.Itoa(server.config.ListeningPort)
	}

	err := endless.ListenAndServe(ListeningPort, Router)
	if err != nil {
		panic(errors.Errorf("serve edge server error %w", err))
	}
}

//HTTP just like a register step,写入/funcName/http/path
func (server *EdgeServer) HTTPGet(path string, handlerFunc gin.HandlerFunc) {

	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj/edgefunc1/path/<path> 是否存在

	Router.GET(path, handlerFunc)
	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/path/<path>
	// v=> 暂时置空,日后留作扩展
}

//HTTP just like a register step,写入/funcName/http/path
func (server *EdgeServer) HTTPPost(path string, handlerFunc gin.HandlerFunc) {
	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj/edgefunc1/path/<path> 是否存在

	Router.POST(path, handlerFunc)

	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/path/<path>
	// v=> 暂时置空,日后留作扩展
}

//WS just like a register step,写入/funcName/ws/path
func (server *EdgeServer) WSGet(path string, handlerFunc gin.HandlerFunc) {

	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj/edgefunc1/path/<path> 是否存在

	Router.GET(path, handlerFunc)

	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/path/<path>
	// v=> 暂时置空,日后留作扩展
}

func Init() *EdgeServer {
	//TODO 此处应该是SDK隐藏,server读取
	return NewForConfigOrDie(&Config{
		MetaConfig: MetaConfig{
			Env:      "dev",
			Org:      "slahser",
			Project:  "ttproj",
			FuncName: "edgefunc1",
		},
		ServerConfig: ServerConfig{
			ListeningPort: 8080,
		},
	})
}
