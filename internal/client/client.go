package client

import (
	"PonGo/internal/pong"
	"log"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
)

type Client struct {
	game *GameOnlinePVP
}

func (c *Client) Update() error {
	c.game.update()
	return nil
}

func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (c *Client) Draw(screen *ebiten.Image) {
	c.game.draw(screen)

}

func Run() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("game client")
	var state pong.GameState

	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Fatal("Couldn’t resolve address:", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()
	state.Reset()

	game := GameOnlinePVP{Game: Game{state: state}, conn: conn}
	go game.pollState()

	if err := ebiten.RunGame(&Client{game: &game}); err != nil {
		log.Fatalln("game engine error :", err)
	}
}
