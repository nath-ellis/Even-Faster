package enemies

import "github.com/hajimehoshi/ebiten/v2"

func Draw(screen *ebiten.Image) {
	for _, e := range Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Obj.X, e.Obj.Y)

		if e.Type == "default-green" {
			screen.DrawImage(enemyCar1, op)
		} else if e.Type == "default-black" {
			screen.DrawImage(enemyCar2, op)
		} else if e.Type == "default-purple" {
			screen.DrawImage(enemyCar3, op)
		} else if e.Type == "default-white" {
			screen.DrawImage(enemyCar4, op)
		}
	}
}
