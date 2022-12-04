package enemies

import (
	"math/rand"

	"github.com/nath-ellis/Even-Faster/player"
	"github.com/solarlune/resolv"
)

// Creates an enemy if needed
func CheckAndCreate() {
	EnemyTimer += 1

	if (EnemyTimer / 60) == SpawnRate {
		newEnemy()
		EnemyTimer = 0
	}
}

func RemoveAll(Space *resolv.Space) {
	EnemyTimer = 0

	objs := Space.Objects()
	for _, o := range objs {
		if o.HasTags("enemy") {
			Space.Remove(o)
		}
	}
	Enemies = make([]Data, 0)
}

func newEnemy() {
	lane1 := rand.Intn(4)
	lane2 := rand.Intn(4)
	lane3 := rand.Intn(4)

	var x int

	if lane1 == 0 {
		x = 155
	} else if lane1 == 1 {
		x = 209
	} else if lane1 == 2 {
		x = 263
	} else {
		x = 317
	}

	ran := rand.Intn(4)
	var typ string

	if ran == 0 {
		typ = "default-green"
	} else if ran == 1 {
		typ = "default-black"
	} else if ran == 2 {
		typ = "default-purple"
	} else if ran == 3 {
		typ = "default-white"
	}

	Enemies = append(Enemies, Data{resolv.NewObject(float64(x), -100, 40, 70, "enemy"), typ, speed()})

	if player.Player.Ticks >= 400 {
		var x int

		if lane2 == 0 {
			x = 155
		} else if lane2 == 1 {
			x = 209
		} else if lane2 == 2 {
			x = 263
		} else {
			x = 317
		}

		ran := rand.Intn(4)
		var typ string

		if ran == 0 {
			typ = "default-green"
		} else if ran == 1 {
			typ = "default-black"
		} else if ran == 2 {
			typ = "default-purple"
		} else if ran == 3 {
			typ = "default-white"
		}

		if lane2 != lane1 {
			Enemies = append(Enemies, Data{resolv.NewObject(float64(x), -100, 40, 70, "enemy"), typ, speed()})
		}
	}

	if player.Player.Ticks >= 800 {
		var x int

		if lane3 == 0 {
			x = 155
		} else if lane3 == 1 {
			x = 209
		} else if lane3 == 2 {
			x = 263
		} else {
			x = 317
		}

		ran := rand.Intn(4)
		var typ string

		if ran == 0 {
			typ = "default-green"
		} else if ran == 1 {
			typ = "default-black"
		} else if ran == 2 {
			typ = "default-purple"
		} else if ran == 3 {
			typ = "default-white"
		}

		if lane2 != lane3 && lane1 != lane3 {
			Enemies = append(Enemies, Data{resolv.NewObject(float64(x), -100, 40, 70, "enemy"), typ, speed()})
		}
	}
}
