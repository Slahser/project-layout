package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/parnurzeal/gorequest"
)

func TestHTTPGet(t *testing.T) {
	request := gorequest.New()
	_, body, _ := request.Get("http://127.0.0.1:8080/tt/useraaaa").End()

	if body != "Hello useraaaa" {
		fmt.Print("not pass")
	}
}

func TestHTTPPost(t *testing.T) {

}

func TestWSGet(t *testing.T) {

	msg := "ping"

	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/ping", nil)
	if err != nil {
		fmt.Errorf("dial: %w", err)
	}
	//defer c.Close()

	werr := c.WriteMessage(websocket.TextMessage, []byte(msg))
	if werr != nil {
		log.Println("write:", werr)
		return
	}
}
