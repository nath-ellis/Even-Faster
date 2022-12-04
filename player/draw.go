package player

import "github.com/hajimehoshi/ebiten/v2"

func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	if Player.SirenOn {
		if Player.SirenStage <= 16 {
			op.GeoM.Translate(-6, 0)
			screen.DrawImage(Player.PlayerCar1, op)
		} else if Player.SirenStage > 16 {
			op.GeoM.Translate(-22, 0)
			screen.DrawImage(Player.PlayerCar2, op)
		}
	} else {
		screen.DrawImage(Player.PlayerCar, op)
	}

	if Exploding {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.2, 1.2)
		op.GeoM.Translate(Player.Obj.X-35, Player.Obj.Y-15)
		if eTicker < 24 {
			screen.DrawImage(explosion[eTicker], op)
			eTicker++
		} else {
			eTicker = 0
			Exploding = false
		}
	}
}
