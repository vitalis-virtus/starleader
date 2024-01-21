package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/assets"
)

type Game struct {
	playerPosition Vector
}

func New() *Game {
	g := Game{
		playerPosition: Vector{
			X: 100,
			Y: 100,
		},
	}
	return &g
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}

	options.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)

	screen.DrawImage(assets.PlayerSprite, &options)

}

func (g *Game) Update() error {
	speed := 2.0 // number of pixels the position changes in a single tick (one Update call)

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.playerPosition.Y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.playerPosition.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.playerPosition.X -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.playerPosition.X += speed
	}

	return nil
}

func (g *Game) Layout(outsideWidth, insideHeight int) (screenWidth, screenHight int) {
	return outsideWidth, insideHeight
}
