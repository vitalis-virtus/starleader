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
	bullets []*Bullet

	meteorSpawnTimer *Timer
}

func New() *Game {
	g := Game{
		meteorSpawnTimer: NewTimer(time.Second),
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
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	return nil
}

// Layout method returns the size of the game window
func (g *Game) Layout(outsideWidth, insideHeight int) (int, int) {
	return screenWidth, screenHeight
}
