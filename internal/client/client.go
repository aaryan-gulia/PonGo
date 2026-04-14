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
	if p == c.state {
		return nil
	}
	c.page.close()
	c.state = p
	switch c.state {
	case MenuPage:
		c.page = &MainMenu{}
	case GameAIPage:
		c.page = &GameAI{}
	case GameOnlinePVPPage:
		c.page = &GameOnlinePVP{}
	case Exit:
		return ebiten.Termination
	}
	c.page.setup()
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

	c := Client{page: &MainMenu{}}
	defer c.page.close()

	if err := ebiten.RunGame(&c); err != nil {
		log.Fatalln("game engine error :", err)
	}
}
