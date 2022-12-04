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
	State string = "menu"
	Space *resolv.Space
	BG    *ebiten.Image
)

func init() {
	Space = resolv.NewSpace(500, 600, 5, 5)

	rand.Seed(time.Now().Unix())

	player.Init(Space)
	enemies.Init()
	ui.InitRoad()
	music.Init()

	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")
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
			music.EndMenuMusic()

			State = "game"
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
			enemies.RemoveAll(Space)

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
		enemies.Draw(screen)

		ui.DrawMenu(screen)
	case "game":
		ui.DrawRoad(screen)

		enemies.Draw(screen)
		player.Draw(screen)

		ui.DrawLives(screen)
		ui.DrawScore(screen)
		ui.DrawInputPrompts(screen)
	case "gameOver":
		enemies.Draw(screen)
		player.Draw(screen)

		ui.DrawGameOver(screen)
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
