package v1

import (
	"github.com/gorilla/websocket"
	"net"
)

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}


