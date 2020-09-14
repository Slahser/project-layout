package edgeserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	ginlimits "github.com/gin-contrib/size"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/configor"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
摘抄一段FC的自定义Server.

监听在任何IP（0.0.0.0）的指定端口（端口可以读取环境变量FC_SERVER_PORT，默认为9000）。
HTTP Server需配置connection keep-alive。
请求超时时间设置为15分钟以上。
HTTP Server需要在25秒内启动完毕。
*/
var (
	Logger   *zap.Logger
	Router   *gin.Engine
	UpGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

)

const (
	ConfigPath    = "./config.toml"
	ListeningPort = ":8080"

	ServerReadTimeOut=180
	ServerWriteTimeOut=180
	ServerMaxHeaderBytes=1 << 20

	LogFile="./logs/edgeserver.log"
	LogMaxSize = 10
	LogMaxBackups = 5
	LogMaxAge =30
	LoggerName="edgeserver"
)

type EdgeServer struct {
	Config Config
}

type Config struct {
	Env     string `required:"true"`
	Org     string `required:"true"`
	Project string `required:"true"`
}

//Init
func Init() *EdgeServer {
	var Config Config

	if err := configor.Load(&Config, ConfigPath); err != nil {
		fmt.Println(errors.Errorf("parse edgeserver config.toml error %w", err))
	}

	return NewForConfigOrDie(&Config)
}

//NewForConfigOrDie
func NewForConfigOrDie(config *Config) *EdgeServer {
	server, err := NewForConfig(config)
	if err != nil {
		fmt.Println(errors.Errorf("init edgeserver error %w", err))
	}
	return server
}

//NewForConfig
func NewForConfig(config *Config) (*EdgeServer, error) {
	c := *config
	server := &EdgeServer{c}
	if err := server.tunning(); err != nil {
		return nil, err
	}
	return server, nil
}

//Start
func (server *EdgeServer) Start() {

	endless.DefaultReadTimeOut = ServerReadTimeOut * time.Second
	endless.DefaultWriteTimeOut = ServerWriteTimeOut * time.Second
	endless.DefaultMaxHeaderBytes = ServerMaxHeaderBytes
	if err := endless.ListenAndServe(ListeningPort, Router); err != nil {
		fmt.Println(errors.Errorf("serve edgeserver error %w", err))
	}
}

//HTTPGet
func (server *EdgeServer) HTTPGet(path string, handlerFunc gin.HandlerFunc) {

	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj/edgefunc1/path/<path> 是否存在

	Router.GET(path, handlerFunc)
	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/path/<path>
	// v=> 暂时置空,日后留作扩展
}

//HTTPPost
func (server *EdgeServer) HTTPPost(path string, handlerFunc gin.HandlerFunc) {
	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj/edgefunc1/path/<path> 是否存在

	Router.POST(path, handlerFunc)

	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/path/<path>
	// v=> 暂时置空,日后留作扩展
}

//WSGet
func (server *EdgeServer) WSGet(path string, handlerFunc gin.HandlerFunc) {

	//TODO 读etcd
	// 检查etcd /dev/slahser/ttproj/edgefunc1/path/<path> 是否存在

	Router.GET(path, handlerFunc)

	//TODO 写入etcd
	// k => /dev/slahser/ttproj/edgefunc1/path/<path>
	// v=> 暂时置空,日后留作扩展
}

//GetWSConn
func GetWSConn(c *gin.Context) *websocket.Conn {
	wsConn, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(errors.Errorf("get websocket connection error %w", err))
	}
	return wsConn
}

func (server *EdgeServer) tunning() error {

	zap.ReplaceGlobals(InitLogger())

	gin.SetMode(gin.DebugMode)

	Router = gin.New()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(ginlimits.RequestSizeLimiter(10))
	Router.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	Router.Use(ginzap.RecoveryWithZap(Logger, true))
	Router.NoRoute(NoRouteHandlerFunc)
	Router.NoMethod(NoMethodHandlerFunc)
	return nil
}

var NoRouteHandlerFunc gin.HandlerFunc = func(context *gin.Context)  {
	context.JSON(http.StatusNotFound, "NoRoute")
}

var NoMethodHandlerFunc gin.HandlerFunc = func(context *gin.Context)  {
	context.JSON(http.StatusMethodNotAllowed,"NoMethod")
}

//InitLogger
func InitLogger() *zap.Logger {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller())
	Logger.Named(LoggerName)
	return Logger
}

//getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

//getLogWriter
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   LogFile,
		MaxSize:    LogMaxSize,
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
