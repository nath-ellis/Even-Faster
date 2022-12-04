package enemies

import (
	"math/rand"

	"github.com/nath-ellis/Even-Faster/music"
	"github.com/nath-ellis/Even-Faster/player"
	"github.com/solarlune/resolv"
)

func Update(Space *resolv.Space) {
	for _, e := range Enemies {
		Space.Add(e.Obj)
		e.Obj.Y += float64(e.Speed)

		if c := e.Obj.Check(0, 0, "enemy"); c != nil {
			pos := c.Objects[0]
			objs := Space.Objects()
			tmp := []Data{}
			for _, o := range objs {
				if pos.X == o.X && pos.Y == o.Y && o.HasTags("enemy") {
					Space.Remove(o)

					for _, E := range Enemies {
						if E.Obj.X == pos.X && E.Obj.Y == pos.Y {
							continue
						}

						if e.Obj.X == E.Obj.X && e.Obj.Y == E.Obj.Y {
							continue
						}

						tmp = append(tmp, E)
					}
				}
			}

			Space.Remove(e.Obj)

			Enemies = []Data{}

			Enemies = tmp
		}

		// On collision with a player
		if c := e.Obj.Check(0, 0, "player"); c != nil {
			music.PlayExplosion()
			player.Exploding = true
			tmp := []Data{}

			for _, E := range Enemies {
				if e.Obj.X == E.Obj.X && e.Obj.Y == E.Obj.Y {
					continue
				}

				tmp = append(tmp, E)
			}

			Enemies = []Data{}
			Enemies = tmp

			Space.Remove(e.Obj)
			player.Player.Lives -= 1
		}

		e.Obj.Update()
	}
}

func speed() int {
	S := rand.Intn(EnemySpeed)

	if S == 0 {
		S += 1
	}

	if S <= player.Player.GameSpeed {
		S += player.Player.GameSpeed
	}

	return S
}

func UpdateSpeed() {
	SpeedTicks += 1

	if (SpeedTicks / 60) == 20 {
		EnemySpeed = 9
	} else if (SpeedTicks / 60) == 40 {
		EnemySpeed = 10
		SpawnRate = 2
	} else if (SpeedTicks / 60) == 60 {
		EnemySpeed = 11
		player.Player.GameSpeed = 3
	} else if (SpeedTicks / 60) == 80 {
		EnemySpeed = 12
		SpawnRate = 1
	} else if (SpeedTicks / 60) == 100 {
		EnemySpeed = 13
	} else if (SpeedTicks / 60) == 120 {
		EnemySpeed = 14
		player.Player.GameSpeed = 4
	}
}
