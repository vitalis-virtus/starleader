package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	player *Player
}

func New() *Game {
	g := Game{
		player: NewPlayer(),
	}
	return &g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

// Layout method returns the size of the game window
func (g *Game) Layout(outsideWidth, insideHeight int) (int, int) {
	return screenWidth, screenHeight
}
