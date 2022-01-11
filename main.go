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
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	State     string = "menu"
	Road      *ebiten.Image
	RoadY1    int = 0
	RoadY2    int = 0
	PlayerCar *ebiten.Image
	Font      font.Face
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

	//PlayerCar, _, _ = ebitenutil.NewImageFromFile("assets/player.png")

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

type Game struct{}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			State = "game"
		}
	case "game":
		updateRoad()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch State {
	case "menu":
		text.Draw(screen, "Click to play", Font, 175, 450, color.White)
	case "game":
		drawRoad(screen)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(50, 135)
		//screen.DrawImage(PlayerCar, op)
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
