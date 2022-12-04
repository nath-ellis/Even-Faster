package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Data struct {
	Obj   *resolv.Object
	Type  string
	Speed int
}

var (
	Enemies    []Data
	enemyCar1  *ebiten.Image
	enemyCar2  *ebiten.Image
	enemyCar3  *ebiten.Image
	enemyCar4  *ebiten.Image
	EnemySpeed int = 8
	EnemyTimer int = 0
	SpawnRate  int = 3
	SpeedTicks int = 1
)

func Init() {
	enemyCar1, _, _ = ebitenutil.NewImageFromFile("assets/enemy1.png")
	enemyCar2, _, _ = ebitenutil.NewImageFromFile("assets/enemy2.png")
	enemyCar3, _, _ = ebitenutil.NewImageFromFile("assets/enemy3.png")
	enemyCar4, _, _ = ebitenutil.NewImageFromFile("assets/enemy4.png")
}
