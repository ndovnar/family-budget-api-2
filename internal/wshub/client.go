package wshub

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	id   string
	conn *websocket.Conn
	send chan any
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		id:   id,
		conn: conn,
		send: make(chan any),
	}
}
