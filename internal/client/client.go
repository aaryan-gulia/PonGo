package client

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type ClientState int

const (
	MenuPage ClientState = iota
	GameAIPage
	GameOnlinePVPPage
	Exit
)

type Page interface {
	update() (ClientState, error)
	draw(screen *ebiten.Image)
	setup()
	close()
}

type Client struct {
	page  Page
	state ClientState
}

func (c *Client) Update() error {
	p, _ := c.page.update()
	if p == Exit {
		return ebiten.Termination
	}
	return nil
}

func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (c *Client) Draw(screen *ebiten.Image) {
	c.page.draw(screen)
}

func Run() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("game client")

	game := GameOnlinePVP{}
	game.setup()
	defer game.close()

	if err := ebiten.RunGame(&Client{page: &game}); err != nil {
		log.Fatalln("game engine error :", err)
	}
}
