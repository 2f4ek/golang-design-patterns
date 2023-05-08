package factory

import "fmt"

type Minion struct {
	attackPower int
	attackRange int
	speed       int
}

func (m *Minion) SetAttackPower(attackPower int) {
	m.attackPower = attackPower
}

func (m *Minion) GetAttackPower() int {
	return m.attackPower
}

func (m *Minion) SetAttackRange(attackRange int) {
	m.attackRange = attackRange
}

func (m *Minion) GetAttackRange() int {
	return m.attackRange
}

func (m *Minion) SetSpeed(speed int) {
	m.speed = speed
}

func (m *Minion) GetSpeed() int {
	return m.speed
}

func (m *Minion) GetDetails() {
	fmt.Printf("Minion: power %d, attack range %d, speed %d \n", m.GetAttackPower(), m.GetAttackRange(), m.GetSpeed())
}
