package player

func Update() {
	if Player.MoveCool > 0 {
		Player.MoveCool -= 1
	}

	if Player.SirenCool > 0 {
		Player.SirenCool -= 1
	}

	if Player.SirenStage > 32 {
		Player.SirenStage = 1
	} else {
		Player.SirenStage += 1
	}
}
