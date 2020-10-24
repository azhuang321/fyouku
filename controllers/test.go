package controllers

import (
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type TestController struct {
	beego.Controller
}

var (
	upgraderDemo = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (c *TestController) Get() {
	c.TplName = "danmuDemo.html"
}

func (c *TestController) WsFunc() {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upgraderDemo.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err = conn.WriteMessage(websocket.TextMessage, []byte("Hello")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		data1 := []byte(string(data)[:] + "======")
		if err = conn.WriteMessage(websocket.TextMessage, data1); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
