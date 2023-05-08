package factory

type SiegeMinion struct {
	Minion
}

func newSiegeMinion() IMinion {
	return &SiegeMinion{
		Minion: Minion{
			attackRange: 12,
			attackPower: 10,
			speed:       3,
		},
	}
}
