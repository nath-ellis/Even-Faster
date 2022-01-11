package main

import (
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
}

var (
	State  string = "menu"
	Road   *ebiten.Image
	RoadY1 int = 0
	RoadY2 int = 0
	Font   font.Face
	Space  *resolv.Space
	player Player
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

	Space = resolv.NewSpace(500, 600, 38, 67)

	player.Obj = resolv.NewObject(155, 400, 38, 67)

	player.PlayerCar1, _, _ = ebitenutil.NewImageFromFile("assets/police1.png")
	player.PlayerCar2, _, _ = ebitenutil.NewImageFromFile("assets/police2.png")
	player.PlayerCar, _, _ = ebitenutil.NewImageFromFile("assets/police.png")

	player.MoveCool = 0

	_, RoadY2 = Road.Size()
	RoadY2 = RoadY2 * -1
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
	screen.DrawImage(player.PlayerCar, op)
}

func movePlayer() {
	// Using ebiten instead of input util because movement is better
	if (ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)) && player.MoveCool <= 0 && player.Obj.X < 317 {
		player.Obj.X += 54
		player.MoveCool += 15
	}

	if (ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)) && player.MoveCool <= 0 && player.Obj.X > 155 {
		player.Obj.X -= 54
		player.MoveCool += 15
	}

	if player.MoveCool > 0 {
		player.MoveCool -= 1
	}
}

type Game struct{}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			State = "game"
		}
	case "game":
		updateRoad()

		movePlayer()
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
