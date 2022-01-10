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
	RoadY     int = 0
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
}

func drawRoad(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(120, float64(RoadY))
	screen.DrawImage(Road, op)
}

type Game struct{}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			State = "game"
		}
	case "game":
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
