package main

import (
)

func main() {
	server := edgeserver.Init()

	server.HTTPGet("/tt/:name", edgeserver.TemplateHTTPGetHandlerFunc)
	server.HTTPPost("/cleanDir", edgeserver.TemplateHTTPPostHandlerFunc)
	server.WSGet("/ping", edgeserver.TemplateWSGetHandlerFunc)

	server.Start()
}
