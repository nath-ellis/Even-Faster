package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Even-Faster/player"
)

var (
	w1, _, _      = ebitenutil.NewImageFromFile("assets/W1.png")
	w2, _, _      = ebitenutil.NewImageFromFile("assets/W2.png")
	a1, _, _      = ebitenutil.NewImageFromFile("assets/A1.png")
	a2, _, _      = ebitenutil.NewImageFromFile("assets/A2.png")
	s1, _, _      = ebitenutil.NewImageFromFile("assets/S1.png")
	s2, _, _      = ebitenutil.NewImageFromFile("assets/S2.png")
	d1, _, _      = ebitenutil.NewImageFromFile("assets/D1.png")
	d2, _, _      = ebitenutil.NewImageFromFile("assets/D2.png")
	f1_1, _, _    = ebitenutil.NewImageFromFile("assets/F1-1.png")
	f1_2, _, _    = ebitenutil.NewImageFromFile("assets/F1-2.png")
	f2_1, _, _    = ebitenutil.NewImageFromFile("assets/F2-1.png")
	f2_2, _, _    = ebitenutil.NewImageFromFile("assets/F2-2.png")
	zero, _, _    = ebitenutil.NewImageFromFile("assets/letters/0.png")
	one, _, _     = ebitenutil.NewImageFromFile("assets/letters/1.png")
	two, _, _     = ebitenutil.NewImageFromFile("assets/letters/2.png")
	three, _, _   = ebitenutil.NewImageFromFile("assets/letters/3.png")
	four, _, _    = ebitenutil.NewImageFromFile("assets/letters/4.png")
	five, _, _    = ebitenutil.NewImageFromFile("assets/letters/5.png")
	six, _, _     = ebitenutil.NewImageFromFile("assets/letters/6.png")
	seven, _, _   = ebitenutil.NewImageFromFile("assets/letters/7.png")
	eight, _, _   = ebitenutil.NewImageFromFile("assets/letters/8.png")
	nine, _, _    = ebitenutil.NewImageFromFile("assets/letters/9.png")
	lifeImg, _, _ = ebitenutil.NewImageFromFile("assets/lives.png")
)

func DrawInputPrompts(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)

	op.GeoM.Translate(35, 535)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		screen.DrawImage(w1, op)
	} else {
		screen.DrawImage(w2, op)
	}

	op.GeoM.Translate(-30, 30)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		screen.DrawImage(a1, op)
	} else {
		screen.DrawImage(a2, op)
	}

	op.GeoM.Translate(30, 0)
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		screen.DrawImage(s1, op)
	} else {
		screen.DrawImage(s2, op)
	}

	op.GeoM.Translate(30, 0)
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		screen.DrawImage(d1, op)
	} else {
		screen.DrawImage(d2, op)
	}

	op.GeoM.Translate(400, -30)
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		screen.DrawImage(f1_1, op)
	} else {
		screen.DrawImage(f1_2, op)
	}

	op.GeoM.Translate(0, 30)
	if ebiten.IsKeyPressed(ebiten.KeyF2) {
		screen.DrawImage(f2_1, op)
	} else {
		screen.DrawImage(f2_2, op)
	}
}

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
			screen.DrawImage(zero, op)
		case "1":
			screen.DrawImage(one, op)
		case "2":
			screen.DrawImage(two, op)
		case "3":
			screen.DrawImage(three, op)
		case "4":
			screen.DrawImage(four, op)
		case "5":
			screen.DrawImage(five, op)
		case "6":
			screen.DrawImage(six, op)
		case "7":
			screen.DrawImage(seven, op)
		case "8":
			screen.DrawImage(eight, op)
		case "9":
			screen.DrawImage(nine, op)
		}
	}
}

func DrawLives(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 43)

	for i := 1; i <= player.Player.Lives; i++ {
		screen.DrawImage(lifeImg, op)
		op.GeoM.Translate(0, 41)
	}
}
