package factory

import "fmt"

func InitMinion(minionType string) (IMinion, error) {
	if minionType == "range" {
		return newRangeMinion(), nil
	}

	if minionType == "melee" {
		return newMeleeMinion(), nil
	}

	if minionType == "siege" {
		return newSiegeMinion(), nil
	}

	return nil, fmt.Errorf("wrong minion type")
}
