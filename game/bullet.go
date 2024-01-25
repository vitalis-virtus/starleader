package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/assets"
)

const (
	bulletSpeedPerSecond = 350.0
)

type Bullet struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

func NewBullet(pos Vector, rotation float64) *Bullet {
	sprite := assets.LaserSprite
	bounds := sprite.Bounds()

	halfX := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	pos.X -= halfX
	pos.Y -= halfY

	return &Bullet{
		position: pos,
		sprite:   sprite,
		rotation: rotation,
	}
}

func (b *Bullet) Update() {
	speed := bulletSpeedPerSecond / float64(ebiten.TPS())

	b.position.X += math.Sin(b.rotation) * speed
	b.position.Y += math.Cos(b.rotation) * -speed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	bounds := b.sprite.Bounds()
	halfX := float64(bounds.Dx()) / 2
	halfY := float64(bounds.Dy()) / 2

	op.GeoM.Translate(-halfX, -halfY)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfX, halfY)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.sprite, op)
}
