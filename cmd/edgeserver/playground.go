package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Slahser/coup-de-grace/pkg/edgeserver"
)

func main() {
	server := edgeserver.Init()

	server.HTTPGet("/tt/:name", TemplateHTTPGetHandlerFunc)
	server.HTTPPost("/cleanDir", TemplateHTTPPostHandlerFunc)
	server.WSGet("/ping", TemplateWSGetHandlerFunc)

	server.Start()
}

//path /user/:name
var TemplateHTTPGetHandlerFunc gin.HandlerFunc = func(context *gin.Context)  {
	name := context.Param("name")
	zap.S().Info("name is : " + name)
	context.String(http.StatusOK, "Hello %s", name)
}

//body name=manu&message=this_is_great
var TemplateHTTPPostHandlerFunc gin.HandlerFunc = func(context *gin.Context) {
	//POSIX
	//sh.Exec(context.Param("dir"))
	name := context.PostForm("name")
	message := context.PostForm("message")
	zap.S().Debug("name is : " + name)
	context.JSON(http.StatusOK, name+message)
}

//path /ping
var TemplateWSGetHandlerFunc gin.HandlerFunc = func(context *gin.Context) {
	wsConn := edgeserver.GetWSConn(context)
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
