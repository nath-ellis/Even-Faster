package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

func Controls(Space *resolv.Space) {
	if (ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && Player.MoveCool <= 0 && Player.Obj.X < 317 {
		Player.Obj.X += 54
		Player.MoveCool += 15
	}

	if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && Player.MoveCool <= 0 && Player.Obj.X > 155 {
		Player.Obj.X -= 54
		Player.MoveCool += 15
	}

	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && Player.Obj.Y > 0 {
		Player.Obj.Y -= 3
	}

	if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && Player.Obj.Y < 533 {
		Player.Obj.Y += 3
	}

	Player.Obj.Update()

	if ebiten.IsKeyPressed(ebiten.KeyF1) && Player.SirenCool <= 0 {
		if Player.SirenOn {
			Player.SirenOn = false
		} else {
			Player.SirenOn = true
		}

		Player.SirenCool += 10
	}
}
