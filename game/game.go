package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/vitalis-virtus/starleader/assets"
)

const (
	screenWidth  = 800
	screenHeight = 600

	meteorSpawnTime     = 1 * time.Second
	baseMeteorVelocity  = 0.25
	meteorSpeedUpAmount = 0.1
	meteorSpeedUpTime   = 5 * time.Second
)

type Game struct {
	player           *Player
	bullets          []*Bullet
	meteors          []*Meteor
	meteorSpawnTimer *Timer

	score int

	baseMeteorVelocity float64
	velocityTimer      *Timer
}

func New() *Game {
	g := Game{
		score:              0,
		meteorSpawnTimer:   NewTimer(meteorSpawnTime),
		baseMeteorVelocity: baseMeteorVelocity,
		velocityTimer:      NewTimer(meteorSpeedUpTime),
	}

	p := NewPlayer(&g)

	g.player = p

	return &g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, screenWidth/2-100, 50, color.White)
}

func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseMeteorVelocity += meteorSpeedUpAmount
	}

	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseMeteorVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	// Check for meteor/bullet collisions
	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
			}
		}
	}

	// Check for meteor/player collisions
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	return nil
}

// Layout method returns the size of the game window
func (g *Game) Layout(outsideWidth, insideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddBullet(bullet *Bullet) {
	g.bullets = append(g.bullets, bullet)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.score = 0
	g.meteors = nil
	g.bullets = nil
}
