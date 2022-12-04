package music

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	explosionSFX *audio.Player
	menuMusic    *audio.Player
	gameMusic    []*audio.Player
	currentTrack int  = 0
	mutedMusic   bool = false
	muteCool     int  = 0
)

func Init() {
	ctx := audio.NewContext(48000)
	f, _ := ebitenutil.OpenFile("assets/explosion.mp3")
	d, _ := mp3.DecodeWithSampleRate(48000, f)
	explosionSFX, _ = ctx.NewPlayer(d)
	explosionSFX.SetVolume(explosionSFX.Volume() + 2)

	f, _ = ebitenutil.OpenFile("assets/music/franticpanic.wav")
	s, _ := wav.DecodeWithSampleRate(48000, f)
	menuMusic, _ = ctx.NewPlayer(s)

	f, _ = ebitenutil.OpenFile("assets/music/cave-disco.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track1, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track1)
	f, _ = ebitenutil.OpenFile("assets/music/cavejam.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track2, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track2)
	f, _ = ebitenutil.OpenFile("assets/music/climber.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track3, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track3)
	f, _ = ebitenutil.OpenFile("assets/music/crusher.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track4, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track4)
	f, _ = ebitenutil.OpenFile("assets/music/cursed-world.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track5, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track5)
	f, _ = ebitenutil.OpenFile("assets/music/dark-woods.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track6, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track6)
	f, _ = ebitenutil.OpenFile("assets/music/mystery-mountain.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track7, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track7)
	f, _ = ebitenutil.OpenFile("assets/music/splash-crashers.wav")
	s, _ = wav.DecodeWithSampleRate(48000, f)
	track8, _ := ctx.NewPlayer(s)
	gameMusic = append(gameMusic, track8)
}
