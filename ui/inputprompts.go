package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	W1, _, _   = ebitenutil.NewImageFromFile("assets/W1.png")
	W2, _, _   = ebitenutil.NewImageFromFile("assets/W2.png")
	A1, _, _   = ebitenutil.NewImageFromFile("assets/A1.png")
	A2, _, _   = ebitenutil.NewImageFromFile("assets/A2.png")
	S1, _, _   = ebitenutil.NewImageFromFile("assets/S1.png")
	S2, _, _   = ebitenutil.NewImageFromFile("assets/S2.png")
	D1, _, _   = ebitenutil.NewImageFromFile("assets/D1.png")
	D2, _, _   = ebitenutil.NewImageFromFile("assets/D2.png")
	F1_1, _, _ = ebitenutil.NewImageFromFile("assets/F1-1.png")
	F1_2, _, _ = ebitenutil.NewImageFromFile("assets/F1-2.png")
	F2_1, _, _ = ebitenutil.NewImageFromFile("assets/F2-1.png")
	F2_2, _, _ = ebitenutil.NewImageFromFile("assets/F2-2.png")
)

func DrawInputPrompts(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)

	op.GeoM.Translate(35, 535)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		screen.DrawImage(W1, op)
	} else {
		screen.DrawImage(W2, op)
	}

	op.GeoM.Translate(-30, 30)
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		screen.DrawImage(A1, op)
	} else {
		screen.DrawImage(A2, op)
	}

	op.GeoM.Translate(30, 0)
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		screen.DrawImage(S1, op)
	} else {
		screen.DrawImage(S2, op)
	}

	op.GeoM.Translate(30, 0)
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		screen.DrawImage(D1, op)
	} else {
		screen.DrawImage(D2, op)
	}

	op.GeoM.Translate(400, -30)
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		screen.DrawImage(F1_1, op)
	} else {
		screen.DrawImage(F1_2, op)
	}

	op.GeoM.Translate(0, 30)
	if ebiten.IsKeyPressed(ebiten.KeyF2) {
		screen.DrawImage(F2_1, op)
	} else {
		screen.DrawImage(F2_2, op)
	}
}
