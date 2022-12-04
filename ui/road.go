package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Even-Faster/player"
)

var (
	Road   *ebiten.Image
	RoadY1 int = 0
	RoadY2 int = 0
)

func InitRoad() {
	Road, _, _ = ebitenutil.NewImageFromFile("assets/road.png")

	_, RoadY2 = Road.Size()
	RoadY2 = RoadY2 * -1
}

func DrawRoad(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(120, float64(RoadY1))
	screen.DrawImage(Road, op)

	op.GeoM.Reset()
	op.GeoM.Translate(120, float64(RoadY2))
	screen.DrawImage(Road, op)
}

func UpdateRoad() {
	RoadY1 += player.Player.GameSpeed
	RoadY2 += player.Player.GameSpeed

	_, roadheight := Road.Size()

	if RoadY1 >= 600 {
		RoadY1 = (roadheight - player.Player.GameSpeed*3) * -1
	}
	if RoadY2 >= 600 {
		RoadY2 = (roadheight - player.Player.GameSpeed*3) * -1
	}
}
