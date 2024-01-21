package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vitalis-virtus/starleader/game"
)

func main() {
	g := game.New()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
