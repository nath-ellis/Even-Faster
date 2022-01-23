package main

import (
	"fmt"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
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
	Obj   *resolv.Object
	Type  string
	Speed int
}

var (
	State       string = "menu"
	Road        *ebiten.Image
	RoadY1      int = 0
	RoadY2      int = 0
	Space       *resolv.Space
	player      Player
	EnemyCar1   *ebiten.Image
	EnemyCar2   *ebiten.Image
	EnemyCar3   *ebiten.Image
	EnemyCar4   *ebiten.Image
	Enemies     []Enemy
	Ticks       int = 0
	EnemyTimer  int = 0
	BG          *ebiten.Image
	Explosion   []*ebiten.Image
	Exploding   bool = false
	ETicker     int  = 0
	AudioPlayer *audio.Player
	Lives       int = 3
	LifeImg     *ebiten.Image
	Score       int = 0
	Zero        *ebiten.Image
	One         *ebiten.Image
	Two         *ebiten.Image
	Three       *ebiten.Image
	Four        *ebiten.Image
	Five        *ebiten.Image
	Six         *ebiten.Image
	Seven       *ebiten.Image
	Eight       *ebiten.Image
	Nine        *ebiten.Image
	Speed       int = 2
	EnemySpeed  int = 8
	SpawnRate   int = 3
	SpeedTicks  int = 1
	Logo        *ebiten.Image
	LeftClick   *ebiten.Image
	GameOver    *ebiten.Image
)

func init() {
	Road, _, _ = ebitenutil.NewImageFromFile("assets/road.png")

	_, RoadY2 = Road.Size()
	RoadY2 = RoadY2 * -1

	Space = resolv.NewSpace(500, 600, 5, 5)

	player.Obj = resolv.NewObject(155, 400, 40, 70, "player")
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

	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")

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

	ctx := audio.NewContext(48000)
	f, _ := ebitenutil.OpenFile("assets/explosion.mp3")
	d, _ := mp3.DecodeWithSampleRate(48000, f)
	AudioPlayer, _ = ctx.NewPlayer(d)

	LifeImg, _, _ = ebitenutil.NewImageFromFile("assets/lives.png")

	Zero, _, _ = ebitenutil.NewImageFromFile("assets/letters/0.png")
	One, _, _ = ebitenutil.NewImageFromFile("assets/letters/1.png")
	Two, _, _ = ebitenutil.NewImageFromFile("assets/letters/2.png")
	Three, _, _ = ebitenutil.NewImageFromFile("assets/letters/3.png")
	Four, _, _ = ebitenutil.NewImageFromFile("assets/letters/4.png")
	Five, _, _ = ebitenutil.NewImageFromFile("assets/letters/5.png")
	Six, _, _ = ebitenutil.NewImageFromFile("assets/letters/6.png")
	Seven, _, _ = ebitenutil.NewImageFromFile("assets/letters/7.png")
	Eight, _, _ = ebitenutil.NewImageFromFile("assets/letters/8.png")
	Nine, _, _ = ebitenutil.NewImageFromFile("assets/letters/9.png")

	Logo, _, _ = ebitenutil.NewImageFromFile("assets/logo.png")
	LeftClick, _, _ = ebitenutil.NewImageFromFile("assets/left-click.png")

	GameOver, _, _ = ebitenutil.NewImageFromFile("assets/game-over.png")
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
	RoadY1 += Speed
	RoadY2 += Speed

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
		player.MoveCool += 20
	}

	if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && player.MoveCool <= 0 && player.Obj.X > 155 {
		player.Obj.X -= 54
		player.MoveCool += 20
	}

	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && player.Obj.Y > 0 {
		player.Obj.Y -= 1
	}

	if (ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && player.Obj.Y < 533 {
		player.Obj.Y += 1
	}

	if c := player.Obj.Check(0, 0, "enemy"); c != nil {
		AudioPlayer.Rewind()
		AudioPlayer.Play()

		Exploding = true

		pos := c.Objects[0]
		objs := Space.Objects()
		tmp := []Enemy{}
		for _, o := range objs {
			if pos.X == o.X && pos.Y == o.Y && o.HasTags("enemy") {
				Space.Remove(o)

				for _, e := range Enemies {
					if e.Obj.X == pos.X && e.Obj.Y == pos.Y {
						continue
					}

					tmp = append(tmp, e)
				}
			}
		}

		Enemies = []Enemy{}

		Enemies = tmp

		Lives -= 1
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
	lane1 := rand.Intn(4)
	lane2 := rand.Intn(4)
	lane3 := rand.Intn(4)

	var x int

	if lane1 == 0 {
		x = 155
	} else if lane1 == 1 {
		x = 209
	} else if lane1 == 2 {
		x = 263
	} else {
		x = 317
	}

	ran := rand.Intn(4)
	var typ string

	if ran == 0 {
		typ = "default-green"
	} else if ran == 1 {
		typ = "default-black"
	} else if ran == 2 {
		typ = "default-purple"
	} else if ran == 3 {
		typ = "default-white"
	}

	Enemies = append(Enemies, Enemy{resolv.NewObject(float64(x), -100, 40, 70, "enemy"), typ, speed()})

	if Ticks >= 400 {
		var x int

		if lane2 == 0 {
			x = 155
		} else if lane2 == 1 {
			x = 209
		} else if lane2 == 2 {
			x = 263
		} else {
			x = 317
		}

		ran := rand.Intn(4)
		var typ string

		if ran == 0 {
			typ = "default-green"
		} else if ran == 1 {
			typ = "default-black"
		} else if ran == 2 {
			typ = "default-purple"
		} else if ran == 3 {
			typ = "default-white"
		}

		if lane2 != lane1 {
			Enemies = append(Enemies, Enemy{resolv.NewObject(float64(x), -100, 40, 70, "enemy"), typ, speed()})
		}
	}

	if Ticks >= 800 {
		var x int

		if lane3 == 0 {
			x = 155
		} else if lane3 == 1 {
			x = 209
		} else if lane3 == 2 {
			x = 263
		} else {
			x = 317
		}

		ran := rand.Intn(4)
		var typ string

		if ran == 0 {
			typ = "default-green"
		} else if ran == 1 {
			typ = "default-black"
		} else if ran == 2 {
			typ = "default-purple"
		} else if ran == 3 {
			typ = "default-white"
		}

		if lane2 != lane3 && lane1 != lane3 {
			Enemies = append(Enemies, Enemy{resolv.NewObject(float64(x), -100, 40, 70, "enemy"), typ, speed()})
		}
	}
}

func moveEnemies() {
	for _, e := range Enemies {
		Space.Add(e.Obj)
		e.Obj.Y += float64(e.Speed)

		if c := e.Obj.Check(0, 0, "enemy"); c != nil {
			pos := c.Objects[0]
			objs := Space.Objects()
			tmp := []Enemy{}
			for _, o := range objs {
				if pos.X == o.X && pos.Y == o.Y && o.HasTags("enemy") {
					Space.Remove(o)

					for _, E := range Enemies {
						if E.Obj.X == pos.X && E.Obj.Y == pos.Y {
							continue
						}

						if e.Obj.X == E.Obj.X && e.Obj.Y == E.Obj.Y {
							continue
						}

						tmp = append(tmp, E)
					}
				}
			}

			Space.Remove(e.Obj)

			Enemies = []Enemy{}

			Enemies = tmp
		}

		e.Obj.Update()
	}
}

func drawEnemies(screen *ebiten.Image) {
	for _, e := range Enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.Obj.X, e.Obj.Y)

		if e.Type == "default-green" {
			screen.DrawImage(EnemyCar1, op)
		} else if e.Type == "default-black" {
			screen.DrawImage(EnemyCar2, op)
		} else if e.Type == "default-purple" {
			screen.DrawImage(EnemyCar3, op)
		} else if e.Type == "default-white" {
			screen.DrawImage(EnemyCar4, op)
		}
	}
}

func speed() int {
	S := rand.Intn(EnemySpeed)

	if S == 0 {
		S += 1
	}

	if S <= Speed {
		S += Speed
	}

	return S
}

func updateSpeed() {
	SpeedTicks += 1

	if (SpeedTicks / 60) == 20 {
		EnemySpeed = 9
	} else if (SpeedTicks / 60) == 40 {
		EnemySpeed = 10
		SpawnRate = 2
	} else if (SpeedTicks / 60) == 60 {
		EnemySpeed = 11
		Speed = 3
	} else if (SpeedTicks / 60) == 80 {
		EnemySpeed = 12
		SpawnRate = 1
	} else if (SpeedTicks / 60) == 100 {
		EnemySpeed = 13
	} else if (SpeedTicks / 60) == 120 {
		EnemySpeed = 14
		Speed = 4
	}
}

func drawScore(screen *ebiten.Image) {
	for i, s := range fmt.Sprint(Score) {
		if i >= len(fmt.Sprint(Score))-1 {
			break
		}

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Scale(0.5, 0.5)
		op.GeoM.Translate(float64((25*i)+5), 5)

		switch string(s) {
		case "0":
			screen.DrawImage(Zero, op)
		case "1":
			screen.DrawImage(One, op)
		case "2":
			screen.DrawImage(Two, op)
		case "3":
			screen.DrawImage(Three, op)
		case "4":
			screen.DrawImage(Four, op)
		case "5":
			screen.DrawImage(Five, op)
		case "6":
			screen.DrawImage(Six, op)
		case "7":
			screen.DrawImage(Seven, op)
		case "8":
			screen.DrawImage(Eight, op)
		case "9":
			screen.DrawImage(Nine, op)
		}
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
		if Lives <= 0 {
			State = "gameOver"
		}

		Ticks += 1
		EnemyTimer += 1

		Score += 1

		updateRoad()

		move()
		updatePlayer()

		if (EnemyTimer / 60) == SpawnRate {
			newEnemy()
			EnemyTimer = 0
		}

		moveEnemies()

		updateSpeed()
	case "gameOver":
		Ticks = 0
		EnemyTimer = 0

		Speed = 2
		EnemySpeed = 8
		SpawnRate = 3
		SpeedTicks = 1

		Score = 0

		updatePlayer()

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			objs := Space.Objects()
			for _, o := range objs {
				if o.HasTags("enemy") {
					Space.Remove(o)
				}
			}
			Enemies = make([]Enemy, 0)

			player.Obj.Y = 400

			Lives = 3

			State = "game"
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(BG, nil)

	switch State {
	case "menu":
		drawRoad(screen)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(98, 100)
		screen.DrawImage(Logo, op)

		op.GeoM.Reset()
		op.GeoM.Scale(4, 4)
		op.GeoM.Translate(225, 400)
		screen.DrawImage(LeftClick, op)
	case "game":
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(5, 43)
		if Lives == 3 {
			screen.DrawImage(LifeImg, op)
			op.GeoM.Translate(0, 41)
			screen.DrawImage(LifeImg, op)
			op.GeoM.Translate(0, 41)
			screen.DrawImage(LifeImg, op)
		} else if Lives == 2 {
			screen.DrawImage(LifeImg, op)
			op.GeoM.Translate(0, 41)
			screen.DrawImage(LifeImg, op)
		} else if Lives == 1 {
			screen.DrawImage(LifeImg, op)
		}

		drawRoad(screen)

		drawEnemies(screen)

		drawPlayer(screen)

		drawScore(screen)
	case "gameOver":
		drawRoad(screen)

		drawEnemies(screen)

		drawPlayer(screen)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(142, 100)
		screen.DrawImage(GameOver, op)

		op.GeoM.Reset()
		op.GeoM.Scale(4, 4)
		op.GeoM.Translate(225, 400)
		screen.DrawImage(LeftClick, op)
	}

	if Exploding {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1.2, 1.2)
		op.GeoM.Translate(player.Obj.X-35, player.Obj.Y-15)
		if ETicker < 24 {
			screen.DrawImage(Explosion[ETicker], op)
			ETicker++
		} else {
			ETicker = 0
			Exploding = false
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 600
}

func main() {
	ebiten.SetWindowSize(500, 600)
	ebiten.SetWindowTitle("Even Faster")

	if err := ebiten.RunGame(&Game{}); err != nil {
		err = nil
	}
}

// ADD COINS / STORE
