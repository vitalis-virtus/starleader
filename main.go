package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, insideHeight int) (screenWidth, screenHight int) {
	return outsideWidth, insideHeight
}

func main() {
	g := Game{}

	err := ebiten.RunGame(&g)
	if err != nil {
		panic(err)
	}
}
