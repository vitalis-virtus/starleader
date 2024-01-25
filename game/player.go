package game

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/assets"
)

const (
	playerRotationPerSecond = math.Pi
	playerShootCooldown     = time.Millisecond * 500

	bulletSpawnOffset = 50.0
)

type Player struct {
	game *Game

	position Vector
	rotation float64
	sprite   *ebiten.Image

	shootCooldown *Timer
}

func NewPlayer(game *Game) *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := Vector{
		X: screenWidth/2 - halfW,
		Y: screenHeight/2 - halfH,
	}

	shootTimer := NewTimer(playerShootCooldown)

	return &Player{
		game:          game,
		position:      position,
		rotation:      0,
		sprite:        sprite,
		shootCooldown: shootTimer,
	}
}

func (p *Player) Update() {
	speed := playerRotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.rotation += speed
	}

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()

		bounds := p.sprite.Bounds()
		halfX := float64(bounds.Dx()) / 2
		halfY := float64(bounds.Dy()) / 2

		bulletSpawnPos := Vector{
			X: p.position.X + halfX + math.Sin(p.rotation)*bulletSpawnOffset,
			Y: p.position.Y + halfY + math.Cos(p.rotation)*-bulletSpawnOffset,
		}

		bullet := NewBullet(bulletSpawnPos, p.rotation)

		p.game.AddBullet(bullet)
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

func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return Rect{
		X:      p.position.X,
		Y:      p.position.Y,
		Width:  float64(bounds.Dx()),
		Height: float64(bounds.Dy()),
	}
}
