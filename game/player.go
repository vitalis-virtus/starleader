package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/assets"
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite
	return &Player{
		position: Vector{
			X: 100,
			Y: 100,
		},
		sprite: sprite,
	}
}

func (p *Player) Update() {
	speed := 2.0 // number of pixels the position changes in a single tick (one Update call)

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.position.Y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.position.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.position.X -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.position.X += speed
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}
