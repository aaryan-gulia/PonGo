package client

import (
	"PonGo/internal/pong"
	"net"
)

type Game struct {
	state pong.GameState
	conn  *net.UDPConn
}
