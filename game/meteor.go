package game

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/assets"
)

type Meteor struct {
	rotation      float64
	rotationSpeed float64

	position Vector
	movement Vector

	sprite *ebiten.Image
}

func NewMeteor() *Meteor {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	target := Vector{
		X: screenWidth / 2,
		Y: screenHeight / 2,
	}

	r := screenWidth / 2.0

	angle := rand.Float64() * 2 * math.Pi

	position := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// Randomized velocity
	velocity := 0.25 + rand.Float64()*1.5

	direction := Vector{
		X: target.X - position.X,
		Y: target.Y - position.Y,
	}

	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	rotationSpeed := -0.02 + rand.Float64()*0.04

	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	return &Meteor{
		position:      position,
		sprite:        sprite,
		movement:      movement,
		rotationSpeed: rotationSpeed,
	}
}

func (m *Meteor) Update() {

	m.rotation += m.rotationSpeed
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y

}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.sprite.Bounds()

	return Rect{
		X:      m.position.X,
		Y:      m.position.Y,
		Width:  float64(bounds.Dx()),
		Height: float64(bounds.Dy()),
	}
}
