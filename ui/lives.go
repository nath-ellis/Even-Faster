package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Even-Faster/player"
)

var lifeImg, _, _ = ebitenutil.NewImageFromFile("assets/lives.png")

func DrawLives(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 43)

	for i := 1; i <= player.Player.Lives; i++ {
		screen.DrawImage(lifeImg, op)
		op.GeoM.Translate(0, 41)
	}
}
