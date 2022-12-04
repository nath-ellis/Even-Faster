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
	EnemyCar1  *ebiten.Image
	EnemyCar2  *ebiten.Image
	EnemyCar3  *ebiten.Image
	EnemyCar4  *ebiten.Image
	EnemySpeed int = 8
	EnemyTimer int = 0
	SpawnRate  int = 3
	SpeedTicks int = 1
)

func Init() {
	EnemyCar1, _, _ = ebitenutil.NewImageFromFile("assets/enemy1.png")
	EnemyCar2, _, _ = ebitenutil.NewImageFromFile("assets/enemy2.png")
	EnemyCar3, _, _ = ebitenutil.NewImageFromFile("assets/enemy3.png")
	EnemyCar4, _, _ = ebitenutil.NewImageFromFile("assets/enemy4.png")
}
