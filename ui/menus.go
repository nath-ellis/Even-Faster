package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	logo, _, _      = ebitenutil.NewImageFromFile("assets/logo.png")
	leftClick, _, _ = ebitenutil.NewImageFromFile("assets/left-click.png")
	gameOver, _, _  = ebitenutil.NewImageFromFile("assets/game-over.png")
)

func DrawMenu(screen *ebiten.Image) {
	DrawRoad(screen)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(98, 100)
	screen.DrawImage(logo, op)

	op.GeoM.Reset()
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(225, 400)
	screen.DrawImage(leftClick, op)
}

func DrawGameOver(screen *ebiten.Image) {
	DrawRoad(screen)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(142, 100)
	screen.DrawImage(gameOver, op)

	op.GeoM.Reset()
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(225, 400)
	screen.DrawImage(leftClick, op)
}
