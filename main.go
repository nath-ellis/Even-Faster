package main

import (
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/solarlune/resolv"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Player struct {
	Obj        *resolv.Object
	PlayerCar1 *ebiten.Image
	PlayerCar2 *ebiten.Image
	PlayerCar  *ebiten.Image
	MoveCool   int
	SirenOn    bool
	SirenStage int
	SirenCool  int
}

type Enemy struct {
	Obj  *resolv.Object
	Type string
}

var (
	State      string = "menu"
	Road       *ebiten.Image
	RoadY1     int = 0
	RoadY2     int = 0
	Font       font.Face
	Space      *resolv.Space
	player     Player
	EnemyCar1  *ebiten.Image
	EnemyCar2  *ebiten.Image
	EnemyCar3  *ebiten.Image
	EnemyCar4  *ebiten.Image
	Enemies    []Enemy
	Ticks      int = 0
	EnemyTimer int = 0
)

func init() {
	Road, _, _ = ebitenutil.NewImageFromFile("assets/road.png")

	b, _ := ioutil.ReadFile("assets/kenney-mini-square.ttf")
	tt, _ := opentype.Parse(b)

	Font, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	_, RoadY2 = Road.Size()
	RoadY2 = RoadY2 * -1

	Space = resolv.NewSpace(500, 600, 38, 67)

	player.Obj = resolv.NewObject(155, 400, 38, 67, "player")
	Space.Add(player.Obj)

	player.PlayerCar1, _, _ = ebitenutil.NewImageFromFile("assets/police1.png")
	player.PlayerCar2, _, _ = ebitenutil.NewImageFromFile("assets/police2.png")
	player.PlayerCar, _, _ = ebitenutil.NewImageFromFile("assets/police.png")

	player.MoveCool = 0

	player.SirenOn = true
	player.SirenStage = 1
	player.SirenCool = 0

	EnemyCar1, _, _ = ebitenutil.NewImageFromFile("assets/enemy1.png")
	EnemyCar2, _, _ = ebitenutil.NewImageFromFile("assets/enemy2.png")
	EnemyCar3, _, _ = ebitenutil.NewImageFromFile("assets/enemy3.png")
	EnemyCar4, _, _ = ebitenutil.NewImageFromFile("assets/enemy4.png")

	rand.Seed(time.Now().Unix())
}

func drawRoad(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(120, float64(RoadY1))
	screen.DrawImage(Road, op)

	op.GeoM.Reset()
	op.GeoM.Translate(120, float64(RoadY2))
	screen.DrawImage(Road, op)
}

func updateRoad() {
	RoadY1 += 1
	RoadY2 += 1

	_, roadheight := Road.Size()

	if RoadY1 >= 600 {
		RoadY1 = -roadheight
	}
	if RoadY2 >= 600 {
		RoadY2 = -roadheight
	}
}

func drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(player.Obj.X, player.Obj.Y)

	if player.SirenOn {
		if player.SirenStage <= 16 {
			op.GeoM.Translate(-6, 0)
			screen.DrawImage(player.PlayerCar1, op)
		} else if player.SirenStage > 16 {
			op.GeoM.Translate(-22, 0)
			screen.DrawImage(player.PlayerCar2, op)
		}
	} else {
		screen.DrawImage(player.PlayerCar, op)
	}
}

func move() {
	// Using ebiten instead of inpututil because movement is better
	if (ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && player.MoveCool <= 0 && player.Obj.X < 317 {
		player.Obj.X += 54
		player.MoveCool += 15
	}

	if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && player.MoveCool <= 0 && player.Obj.X > 155 {
		player.Obj.X -= 54
		player.MoveCool += 15
	}

	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && player.Obj.Y > 0 {
		player.Obj.Y -= 1
	}

	if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && player.Obj.Y < 533 {
		player.Obj.Y += 1
	}

	if c := player.Obj.Check(0, 0); c != nil {
		State = "gameOver"
	}

	player.Obj.Update()
}

func updatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyZ) && player.SirenCool <= 0 {
		if player.SirenOn {
			player.SirenOn = false
		} else {
			player.SirenOn = true
		}

		player.SirenCool += 10
	}

	if player.MoveCool > 0 {
		player.MoveCool -= 1
	}

	if player.SirenCool > 0 {
		player.SirenCool -= 1
	}

	if player.SirenStage > 32 {
		player.SirenStage = 1
	} else {
		player.SirenStage += 1
	}
}

func newEnemy() {
	lane := rand.Intn(4)

	var x int

	if lane == 0 {
		x = 155
	} else if lane == 1 {
		x = 209
	} else if lane == 2 {
		x = 263
	} else {
		x = 317
	}

	Enemies = append(Enemies, Enemy{resolv.NewObject(float64(x), -100, 38, 67, "enemy"), "default-green"})
}

func moveEnemies() {
	for _, e := range Enemies {
		e.Obj.Y += 2
		e.Obj.Update()
	}
}

func drawEnemies(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, e := range Enemies {
		op.GeoM.Translate(e.Obj.X, e.Obj.Y)
		screen.DrawImage(EnemyCar1, op)
		op.GeoM.Reset()
	}
}

type Game struct{}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			State = "game"
		}
	case "game":
		Ticks += 1

		EnemyTimer += 1

		updateRoad()

		move()
		updatePlayer()

		if (EnemyTimer / 60) == 3 {
			newEnemy()
			EnemyTimer = 0
		}

		moveEnemies()
	case "gameOver":
		Ticks = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch State {
	case "menu":
		text.Draw(screen, "Click to play", Font, 175, 450, color.White)
	case "game":
		drawRoad(screen)

		drawPlayer(screen)

		drawEnemies(screen)
	case "gameOver":
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 600
}

func main() {
	ebiten.SetWindowSize(500, 600)
	ebiten.SetWindowTitle("Even Faster")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal("Failed to Run: ", err)
	}
}
