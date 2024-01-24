package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/assets"
)

type Player struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := Vector{
		X: screenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	return &Player{
		position: position,
		rotation: 0,
		sprite:   sprite,
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.rotation += speed
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
