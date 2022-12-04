package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Data struct {
	Obj        *resolv.Object
	PlayerCar1 *ebiten.Image
	PlayerCar2 *ebiten.Image
	PlayerCar  *ebiten.Image
	MoveCool   int
	SirenOn    bool
	SirenStage int
	SirenCool  int
	Score      int
	Lives      int
	Ticks      int
	GameSpeed  int
}

var (
	Player    Data
	Explosion []*ebiten.Image
	Exploding bool = false
	ETicker   int  = 0
)

func Init(Space *resolv.Space) {
	Player.Obj = resolv.NewObject(155, 400, 40, 70, "player")
	Space.Add(Player.Obj)

	Player.PlayerCar1, _, _ = ebitenutil.NewImageFromFile("assets/police1.png")
	Player.PlayerCar2, _, _ = ebitenutil.NewImageFromFile("assets/police2.png")
	Player.PlayerCar, _, _ = ebitenutil.NewImageFromFile("assets/police.png")

	Player.MoveCool = 0

	Player.SirenOn = true
	Player.SirenStage = 1
	Player.SirenCool = 0

	Player.Score = 0
	Player.Lives = 3
	Player.Ticks = 0
	Player.GameSpeed = 2

	explosion1, _, _ := ebitenutil.NewImageFromFile("assets/explosion1.png")
	explosion2, _, _ := ebitenutil.NewImageFromFile("assets/explosion2.png")
	explosion3, _, _ := ebitenutil.NewImageFromFile("assets/explosion3.png")
	explosion4, _, _ := ebitenutil.NewImageFromFile("assets/explosion4.png")
	explosion5, _, _ := ebitenutil.NewImageFromFile("assets/explosion5.png")
	explosion6, _, _ := ebitenutil.NewImageFromFile("assets/explosion6.png")
	explosion7, _, _ := ebitenutil.NewImageFromFile("assets/explosion7.png")
	explosion8, _, _ := ebitenutil.NewImageFromFile("assets/explosion8.png")
	explosion9, _, _ := ebitenutil.NewImageFromFile("assets/explosion9.png")
	explosion10, _, _ := ebitenutil.NewImageFromFile("assets/explosion10.png")
	explosion11, _, _ := ebitenutil.NewImageFromFile("assets/explosion11.png")
	explosion12, _, _ := ebitenutil.NewImageFromFile("assets/explosion12.png")

	Explosion = append(Explosion, explosion1)
	Explosion = append(Explosion, explosion1)
	Explosion = append(Explosion, explosion2)
	Explosion = append(Explosion, explosion2)
	Explosion = append(Explosion, explosion3)
	Explosion = append(Explosion, explosion3)
	Explosion = append(Explosion, explosion4)
	Explosion = append(Explosion, explosion4)
	Explosion = append(Explosion, explosion5)
	Explosion = append(Explosion, explosion5)
	Explosion = append(Explosion, explosion6)
	Explosion = append(Explosion, explosion6)
	Explosion = append(Explosion, explosion7)
	Explosion = append(Explosion, explosion7)
	Explosion = append(Explosion, explosion8)
	Explosion = append(Explosion, explosion8)
	Explosion = append(Explosion, explosion9)
	Explosion = append(Explosion, explosion9)
	Explosion = append(Explosion, explosion10)
	Explosion = append(Explosion, explosion10)
	Explosion = append(Explosion, explosion11)
	Explosion = append(Explosion, explosion11)
	Explosion = append(Explosion, explosion12)
	Explosion = append(Explosion, explosion12)
}
