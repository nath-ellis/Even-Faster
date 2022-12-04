package music

func Update() {
	playing := false

	for _, g := range gameMusic {
		if g.IsPlaying() {
			playing = true
		}
	}

	if !playing {
		gameMusic[currentTrack].Rewind()
		gameMusic[currentTrack].Play()

		currentTrack += 1
	}

	if currentTrack == (len(gameMusic) - 1) {
		currentTrack = 0
	}
}
