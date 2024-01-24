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

func (g *Game) Layout(outsideWidth, insideHeight int) (screenWidth, screenHight int) {
	return screenWidth, screenHeight
}
