package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	player  *Player
	meteors []*Meteor
	bullet  *Bullet

	meteorSpawnTimer *Timer
}

func New() *Game {
	g := Game{
		player:           NewPlayer(),
		meteorSpawnTimer: NewTimer(time.Second),
		bullet: NewBullet(Vector{
			screenWidth / 2, screenHeight / 2,
		}, 45.0),
	}
	return &g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	g.bullet.Draw(screen)
	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Update() error {
	g.player.Update()
	g.bullet.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	return nil
}

// Layout method returns the size of the game window
func (g *Game) Layout(outsideWidth, insideHeight int) (int, int) {
	return screenWidth, screenHeight
}
