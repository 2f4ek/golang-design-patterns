package factory

type RangeMinion struct {
	Minion
}

func newRangeMinion() IMinion {
	return &RangeMinion{
		Minion: Minion{
			attackRange: 10,
			attackPower: 5,
			speed:       5,
		},
	}
}
