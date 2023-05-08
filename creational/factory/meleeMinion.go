package factory

type MeleeMinion struct {
	Minion
}

func newMeleeMinion() IMinion {
	return &MeleeMinion{
		Minion: Minion{
			attackRange: 1,
			attackPower: 6,
			speed:       6,
		},
	}
}
