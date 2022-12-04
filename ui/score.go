package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Even-Faster/player"
)

var (
	Zero, _, _  = ebitenutil.NewImageFromFile("assets/letters/0.png")
	One, _, _   = ebitenutil.NewImageFromFile("assets/letters/1.png")
	Two, _, _   = ebitenutil.NewImageFromFile("assets/letters/2.png")
	Three, _, _ = ebitenutil.NewImageFromFile("assets/letters/3.png")
	Four, _, _  = ebitenutil.NewImageFromFile("assets/letters/4.png")
	Five, _, _  = ebitenutil.NewImageFromFile("assets/letters/5.png")
	Six, _, _   = ebitenutil.NewImageFromFile("assets/letters/6.png")
	Seven, _, _ = ebitenutil.NewImageFromFile("assets/letters/7.png")
	Eight, _, _ = ebitenutil.NewImageFromFile("assets/letters/8.png")
	Nine, _, _  = ebitenutil.NewImageFromFile("assets/letters/9.png")
)

func DrawScore(screen *ebiten.Image) {
	for i, s := range fmt.Sprint(player.Player.Score) {
		if i >= len(fmt.Sprint(player.Player.Score))-1 {
			break
		}

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Scale(0.5, 0.5)
		op.GeoM.Translate(float64((25*i)+5), 5)

		switch string(s) {
		case "0":
			screen.DrawImage(Zero, op)
		case "1":
			screen.DrawImage(One, op)
		case "2":
			screen.DrawImage(Two, op)
		case "3":
			screen.DrawImage(Three, op)
		case "4":
			screen.DrawImage(Four, op)
		case "5":
			screen.DrawImage(Five, op)
		case "6":
			screen.DrawImage(Six, op)
		case "7":
			screen.DrawImage(Seven, op)
		case "8":
			screen.DrawImage(Eight, op)
		case "9":
			screen.DrawImage(Nine, op)
		}
	}
}
