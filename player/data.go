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
	explosion []*ebiten.Image
	Exploding bool = false
	eTicker   int  = 0
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

	// Each frame twice to extend the animation
	explosion = append(explosion, explosion1)
	explosion = append(explosion, explosion1)
	explosion = append(explosion, explosion2)
	explosion = append(explosion, explosion2)
	explosion = append(explosion, explosion3)
	explosion = append(explosion, explosion3)
	explosion = append(explosion, explosion4)
	explosion = append(explosion, explosion4)
	explosion = append(explosion, explosion5)
	explosion = append(explosion, explosion5)
	explosion = append(explosion, explosion6)
	explosion = append(explosion, explosion6)
	explosion = append(explosion, explosion7)
	explosion = append(explosion, explosion7)
	explosion = append(explosion, explosion8)
	explosion = append(explosion, explosion8)
	explosion = append(explosion, explosion9)
	explosion = append(explosion, explosion9)
	explosion = append(explosion, explosion10)
	explosion = append(explosion, explosion10)
	explosion = append(explosion, explosion11)
	explosion = append(explosion, explosion11)
	explosion = append(explosion, explosion12)
	explosion = append(explosion, explosion12)
}
