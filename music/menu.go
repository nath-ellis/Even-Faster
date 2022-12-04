package music

func RestartMenuMusic() {
	if !menuMusic.IsPlaying() {
		menuMusic.Rewind()
		menuMusic.Play()
	}
}

func EndMenuMusic() {
	menuMusic.Pause()
	menuMusic.Close()
}
