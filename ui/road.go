package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Even-Faster/player"
)

var (
	road   *ebiten.Image
	roadY1 int = 0
	roadY2 int = 0
)

func InitRoad() {
	road, _, _ = ebitenutil.NewImageFromFile("assets/road.png")

	_, roadY2 = road.Size()
	roadY2 = roadY2 * -1
}

func DrawRoad(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(120, float64(roadY1))
	screen.DrawImage(road, op)

	op.GeoM.Reset()
	op.GeoM.Translate(120, float64(roadY2))
	screen.DrawImage(road, op)
}

func UpdateRoad() {
	roadY1 += player.Player.GameSpeed
	roadY2 += player.Player.GameSpeed

	_, roadheight := road.Size()

	if roadY1 >= 600 {
		roadY1 = (roadheight - player.Player.GameSpeed*3) * -1
	}
	if roadY2 >= 600 {
		roadY2 = (roadheight - player.Player.GameSpeed*3) * -1
	}
}
