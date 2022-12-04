package main

import (
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Even-Faster/enemies"
	"github.com/nath-ellis/Even-Faster/music"
	"github.com/nath-ellis/Even-Faster/player"
	"github.com/nath-ellis/Even-Faster/ui"
	"github.com/solarlune/resolv"
)

var (
	State        string = "menu"
	Space        *resolv.Space
	BG           *ebiten.Image
	Logo         *ebiten.Image
	LeftClick    *ebiten.Image
	GameOver     *ebiten.Image
	SeenControls bool = false
)

func init() {
	Space = resolv.NewSpace(500, 600, 5, 5)

	rand.Seed(time.Now().Unix())

	player.Init(Space)
	enemies.Init()
	ui.InitRoad()

	music.Init()

	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")

	Logo, _, _ = ebitenutil.NewImageFromFile("assets/logo.png")
	LeftClick, _, _ = ebitenutil.NewImageFromFile("assets/left-click.png")

	GameOver, _, _ = ebitenutil.NewImageFromFile("assets/game-over.png")
}

type Game struct{}

func (g *Game) Update() error {
	switch State {
	case "menu":
		music.RestartMenuMusic()

		enemies.CheckAndCreate()
		enemies.Update(Space)

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

			enemies.RemoveAll(Space)

			State = "game"

			music.EndMenuMusic()
		}
	case "game":
		if player.Player.Lives <= 0 {
			State = "gameOver"
		}

		player.Player.Ticks += 1
		enemies.EnemyTimer += 1

		player.Player.Score += 1

		ui.UpdateRoad()

		player.Controls(Space)
		player.Update()

		music.ControlGameMusic()

		enemies.CheckAndCreate()
		enemies.Update(Space)
		enemies.UpdateSpeed()
	case "gameOver":
		music.EndGameMusic()

		player.Player.Ticks = 0
		enemies.EnemyTimer = 0

		player.Player.GameSpeed = 2
		enemies.EnemySpeed = 8
		enemies.SpawnRate = 3
		enemies.SpeedTicks = 1

		player.Player.Score = 0

		player.Update()

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			objs := Space.Objects()
			for _, o := range objs {
				if o.HasTags("enemy") {
					Space.Remove(o)
				}
			}
			enemies.Enemies = make([]enemies.Data, 0)

			player.Player.Obj.Y = 400

			player.Player.Lives = 3

			State = "game"
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(BG, nil)

	switch State {
	case "menu":
		ui.DrawRoad(screen)

		enemies.Draw(screen)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(98, 100)
		screen.DrawImage(Logo, op)

		op.GeoM.Reset()
		op.GeoM.Scale(4, 4)
		op.GeoM.Translate(225, 400)
		screen.DrawImage(LeftClick, op)
	case "game":
		ui.DrawLives(screen)
		ui.DrawRoad(screen)
		enemies.Draw(screen)
		player.Draw(screen)
		ui.DrawScore(screen)

		if !SeenControls {
			ui.DrawInputPrompts(screen)

			if (player.Player.Ticks / 60) >= 10 {
				SeenControls = true
			}
		}
	case "gameOver":
		ui.DrawRoad(screen)
		enemies.Draw(screen)
		player.Draw(screen)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(142, 100)
		screen.DrawImage(GameOver, op)

		op.GeoM.Reset()
		op.GeoM.Scale(4, 4)
		op.GeoM.Translate(225, 400)
		screen.DrawImage(LeftClick, op)
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
