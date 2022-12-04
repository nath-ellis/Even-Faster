package music

import "github.com/hajimehoshi/ebiten/v2"

func ControlGameMusic() {
	if ebiten.IsKeyPressed(ebiten.KeyF2) && muteCool <= 0 {
		if mutedMusic {
			mutedMusic = false
		} else {
			mutedMusic = true
		}

		muteCool = 10
	} else {
		if muteCool != 0 {
			muteCool -= 1
		}
	}

	if !mutedMusic {
		Update()
	} else {
		EndGameMusic()
	}
}

func EndGameMusic() {
	for _, g := range gameMusic {
		if g.IsPlaying() {
			g.Pause()
		}
	}
}
