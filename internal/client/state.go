package client

import (
	"PonGo/internal/pong"
	"net"
)

type ConnectionState int
type GameType int
type ClientState int

const (
	Connected ConnectionState = iota
	Waiting
	Paused
	Playing
	Disconnected
)

const (
	Ai GameType = iota
	OnlinePvp
)

const (
	MainMenuPage ClientState = iota
	GamePage
)

type Game struct {
	state pong.GameState
	conn  *net.UDPConn
}

type Client struct {
	state ClientState

	game *Game
}
